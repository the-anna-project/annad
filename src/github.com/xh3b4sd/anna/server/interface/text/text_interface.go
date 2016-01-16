package textinterface

import (
	"fmt"
	"time"

	"github.com/xh3b4sd/anna/gateway"
)

type TextInterface interface {
	FetchURL(url string) ([]byte, error)
	ReadFile(file string) ([]byte, error)
	ReadStream(stream string) ([]byte, error)
	ReadPlain(plain []byte) ([]byte, error)
}

type TextInterfaceConfig struct {
	TextGateway gateway.Gateway
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

func (ti textInterface) ReadPlain(plain []byte) ([]byte, error) {
	// TODO we need a task system here to make the calls asynchronous
	newSignal := gateway.NewSignal(plain)

	i := 0
	for {
		if i >= 5 {
			return nil, maskAny(gatewayClosedError)
		}
		i++

		err := ti.TextGateway.SendSignal(newSignal)
		if gateway.IsGatewayClosed(err) {
			fmt.Printf("gateway is closed\n")
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			fmt.Printf("%#v\n", maskAny(err))
		} else {
			break
		}
	}

	return <-newSignal.GetResponder(), nil
}
