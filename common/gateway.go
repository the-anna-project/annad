package common

import (
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/gateway/spec"
)

func ForwardSignal(gw gatewayspec.Gateway, bytes map[string][]byte) (interface{}, error) {
	newConfig := gateway.DefaultSignalConfig()
	for key, val := range bytes {
		newConfig.Bytes[key] = val
	}
	newSignal := gateway.NewSignal(newConfig)

	err := gw.SendSignal(newSignal)
	if err != nil {
		return nil, maskAny(err)
	}

	responder, err := newSignal.GetResponder()
	if err != nil {
		return nil, maskAny(err)
	}

	resSignal := <-responder
	response, err := resSignal.GetObject("response")
	if err != nil {
		return nil, maskAny(err)
	}

	return response, nil
}
