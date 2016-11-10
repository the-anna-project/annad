package main

import (
	"github.com/xh3b4sd/anna/server/interface/text"
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// TODO text interface should be a service inside the service collection
func newTextInterface(newServiceCollection servicespec.Collection) (text.TextInterfaceServer, error) {
	newTextInterfaceConfig := text.DefaultServerConfig()
	newTextInterfaceConfig.ServiceCollection = newServiceCollection
	newTextInterface, err := text.NewServer(newTextInterfaceConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newTextInterface, nil
}
