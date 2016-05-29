package main

import (
	"github.com/xh3b4sd/anna/server/interface/text"
	"github.com/xh3b4sd/anna/spec"
)

func createTextInterface(newLog spec.Log, newScheduler spec.Scheduler, newTextGateway spec.Gateway) (spec.TextInterface, error) {
	newTextInterfaceConfig := text.DefaultInterfaceConfig()
	newTextInterfaceConfig.Log = newLog
	newTextInterfaceConfig.Scheduler = newScheduler
	newTextInterfaceConfig.TextGateway = newTextGateway
	newTextInterface, err := text.NewInterface(newTextInterfaceConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newTextInterface, nil
}
