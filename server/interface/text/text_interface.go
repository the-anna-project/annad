package textinterface

import (
	"time"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/gateway/spec"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/server/spec"
	"github.com/xh3b4sd/anna/spec"
)

type Config struct {
	Log spec.Log

	TextGateway gatewayspec.Gateway
}

func DefaultConfig() Config {
	return Config{
		Log:         log.NewLog(log.DefaultConfig()),
		TextGateway: gateway.NewGateway(),
	}
}

func NewTextInterface(config Config) serverspec.TextInterface {
	return textInterface{
		Config: config,
	}
}

type textInterface struct {
	Config
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
	ti.Log.V(11).Debugf("call TextInterface.ReadPlainWithID")

	newConfig := gateway.DefaultSignalConfig()
	newConfig.Bytes["request"] = []byte{}
	newConfig.ID = ID
	newSignal := gateway.NewSignal(newConfig)

	response, err := ti.waitForSignal(ctx, newSignal)
	if err != nil {
		return "", maskAny(err)
	}

	return response, nil
}

// return ID
func (ti textInterface) ReadPlainWithPlain(ctx context.Context, plain string) (string, error) {
	ti.Log.V(11).Debugf("call TextInterface.ReadPlainWithPlain")

	newConfig := gateway.DefaultSignalConfig()
	newConfig.Bytes["request"] = []byte(plain)
	newSignal := gateway.NewSignal(newConfig)

	response, err := ti.waitForSignal(ctx, newSignal)
	if err != nil {
		return "", maskAny(err)
	}

	return response, nil
}

func (ti textInterface) waitForSignal(ctx context.Context, signal gatewayspec.Signal) (string, error) {
	ti.Log.V(11).Debugf("call TextInterface.waitForSignal")

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
	ti.Log.V(11).Debugf("call TextInterface.forwardSignal")

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
			ti.Log.V(6).Warnf("gateway is closed")
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
