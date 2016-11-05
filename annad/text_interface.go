package main

import (
	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/server/interface/text"
	"github.com/xh3b4sd/anna/spec"
)

// TODO text interface should be a service inside the service collection
func newTextInterface(newLog spec.Log, newServiceCollection spec.ServiceCollection, newTextInput chan spec.TextRequest) (api.TextInterfaceServer, error) {
	newTextInterfaceConfig := text.DefaultServerConfig()
	newTextInterfaceConfig.Log = newLog
	newTextInterfaceConfig.TextInput = newTextInput
	newTextInterfaceConfig.ServiceCollection = newServiceCollection
	newTextInterface, err := text.NewServer(newTextInterfaceConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newTextInterface, nil
}
