package main

import (
	"github.com/xh3b4sd/anna/client/interface/text"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
)

// TODO text interface should be a service inside the service collection
func newTextInterface(newServiceCollection servicespec.Collection, gRPCAddr string) (systemspec.TextInterfaceClient, error) {
	textInterfaceConfig := text.DefaultClientConfig()
	textInterfaceConfig.GRPCAddr = gRPCAddr
	textInterfaceConfig.ServiceCollection = newServiceCollection
	newClient, err := text.NewClient(textInterfaceConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newClient, nil
}
