package main

import (
	"github.com/xh3b4sd/anna/server/interface/text"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
)

// TODO text interface should be a service inside the service collection
func newTextInterface(newLog systemspec.Log, newServiceCollection servicespec.Collection) (text.TextInterfaceServer, error) {
	newTextInterfaceConfig := text.DefaultServerConfig()
	newTextInterfaceConfig.Log = newLog
	newTextInterfaceConfig.ServiceCollection = newServiceCollection
	newTextInterface, err := text.NewServer(newTextInterfaceConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newTextInterface, nil
}
