package main

import (
	"github.com/xh3b4sd/anna/client/interface/text"
	"github.com/xh3b4sd/anna/spec"
)

// TODO text interface should be a service inside the service collection
func newTextInterface(newServiceCollection spec.ServiceCollection, gRPCAddr string) (spec.TextInterfaceClient, error) {
	textInterfaceConfig := text.DefaultClientConfig()
	textInterfaceConfig.GRPCAddr = gRPCAddr
	textInterfaceConfig.ServiceCollection = newServiceCollection
	newClient, err := text.NewClient(textInterfaceConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newClient, nil
}
