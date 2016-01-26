package textinterface

import (
	"fmt"
	"time"

	"github.com/xh3b4sd/anna/gateway"
	gatewayspec "github.com/xh3b4sd/anna/gateway/spec"
	serverspec "github.com/xh3b4sd/anna/server/spec"
	"golang.org/x/net/context"
)

type TextInterfaceConfig struct {
	TextGateway gatewayspec.Gateway
}

func DefaultTextInterfaceConfig() TextInterfaceConfig {
	return TextInterfaceConfig{
		TextGateway: nil,
	}
}

func NewTextInterface(config TextInterfaceConfig) serverspec.TextInterface {
	return textInterface{
		TextInterfaceConfig: config,
	}
}

type textInterface struct {
	TextInterfaceConfig
}

// TODO this should actually fetch a url from the web
func (ti textInterface) FetchURL(url string) ([]byte, error) {
	return nil, nil
}

// TODO this should actually read a file from file system
func (ti textInterface) ReadFile(file string) ([]byte, error) {
	return nil, nil
}

// TODO this should actually be streamed
func (ti textInterface) ReadStream(stream string) ([]byte, error) {
	return nil, nil
}

// return response
func (ti textInterface) ReadPlainWithID(ctx context.Context, ID string) (string, error) {
	newSignalConfig := gateway.DefaultSignalConfig()
	newSignalConfig.ID = ID
	newSignal := gateway.NewSignal(newSignalConfig)

	response, err := ti.waitForSignal(ctx, newSignal)
	if err != nil {
		return "", maskAny(err)
	}

	return response, nil
}

// return ID
func (ti textInterface) ReadPlainWithPlain(ctx context.Context, plain string) (string, error) {
	newSignalConfig := gateway.DefaultSignalConfig()
	newSignalConfig.Bytes["request"] = []byte(plain)
	newSignal := gateway.NewSignal(newSignalConfig)

	response, err := ti.waitForSignal(ctx, newSignal)
	if err != nil {
		return "", maskAny(err)
	}

	return response, nil
}

func (ti textInterface) waitForSignal(ctx context.Context, signal gatewayspec.Signal) (string, error) {
	for {
		response, err := ti.forwardSignal(ctx, signal)
		if err != nil {
			return "", maskAny(err)
		}
		if len(response) == 0 {
			time.Sleep(1 * time.Second)
		} else {
			return response, nil
		}
	}
}

func (ti textInterface) forwardSignal(ctx context.Context, signal gatewayspec.Signal) (string, error) {
	var err error
	var response []byte

	i := 0
	for {
		if i >= 5 {
			return "", maskAny(gatewayClosedError)
		}

		err := ti.TextGateway.SendSignal(signal)
		if gateway.IsGatewayClosed(err) {
			i++
			fmt.Printf("gateway is closed\n")
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			return "", maskAny(err)
		}

		// Once we send the signal through the gateway, we stop here, to prevent
		// resubmition.
		break
	}

	responder, err := signal.GetResponder()
	if err != nil {
		return "", maskAny(err)
	}

	select {
	case <-ctx.Done():
		signal.Cancel()
		return "", nil
	case resSignal := <-responder:
		if err := resSignal.GetError(); err != nil {
			return "", maskAny(err)
		}

		response, err = resSignal.GetBytes("response")
		if err != nil {
			return "", maskAny(err)
		}
	}

	return string(response), nil
}
