package factoryclient

import (
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/spec"
)

func forwardSignal(g spec.Gateway, input interface{}, closer <-chan struct{}) (interface{}, error) {
	newConfig := gateway.DefaultSignalConfig()
	newConfig.Input = input
	newSignal := gateway.NewSignal(newConfig)

	var err error

	newSignal, err = g.Send(newSignal, closer)
	if err != nil {
		return nil, maskAny(err)
	}

	output := newSignal.GetOutput()

	return output, nil
}
