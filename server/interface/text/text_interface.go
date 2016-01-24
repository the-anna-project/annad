package textinterface

import (
	"fmt"
	"time"

	"github.com/xh3b4sd/anna/gateway"
	gatewayspec "github.com/xh3b4sd/anna/gateway/spec"
	"golang.org/x/net/context"
)

type TextInterface interface {
	FetchURL(url string) ([]byte, error)
	ReadFile(file string) ([]byte, error)
	ReadStream(stream string) ([]byte, error)
	ReadPlainWithID(ctx context.Context, id string) ([]byte, error)
	ReadPlainWithPlain(ctx context.Context, plain []byte) (string, error)
}

type TextInterfaceConfig struct {
	TextGateway gatewayspec.Gateway
}

func DefaultTextInterfaceConfig() TextInterfaceConfig {
	return TextInterfaceConfig{
		TextGateway: nil,
	}
}

func NewTextInterface(config TextInterfaceConfig) TextInterface {
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
func (ti textInterface) ReadPlainWithID(ctx context.Context, id string) ([]byte, error) {
	newSignalConfig := gateway.DefaultSignalConfig()
	newSignalConfig.ID = id
	newSignal := gateway.NewSignal(newSignalConfig)

	response, err := ti.forwardSignal(ctx, newSignal)
	if err != nil {
		return nil, maskAny(err)
	}

	return response, nil
}

// return ID
func (ti textInterface) ReadPlainWithPlain(ctx context.Context, plain []byte) (string, error) {
	newSignalConfig := gateway.DefaultSignalConfig()
	newSignalConfig.Bytes["request"] = plain
	newSignal := gateway.NewSignal(newSignalConfig)

	response, err := ti.forwardSignal(ctx, newSignal)
	if err != nil {
		return "", maskAny(err)
	}

	return string(response), nil
}

func (ti textInterface) forwardSignal(ctx context.Context, signal gatewayspec.Signal) ([]byte, error) {
	var err error
	var response []byte

	i := 0
	for {
		if i >= 5 {
			return nil, maskAny(gatewayClosedError)
		}

		err := ti.TextGateway.SendSignal(signal)
		if gateway.IsGatewayClosed(err) {
			i++
			fmt.Printf("gateway is closed\n")
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			return nil, maskAny(err)
		}

		// Once we send the signal through the gateway, we stop here, to prevent
		// resubmition.
		break
	}

	responder, err := signal.GetResponder()
	if err != nil {
		return nil, maskAny(err)
	}

	select {
	case <-ctx.Done():
		signal.Cancel()
		return nil, nil
	case resSignal := <-responder:
		if err := resSignal.GetError(); err != nil {
			return nil, maskAny(err)
		}

		response, err = resSignal.GetBytes("response")
		if err != nil {
			return nil, maskAny(err)
		}
	}

	return response, nil
}
