package main

import (
	systemspec "github.com/the-anna-project/spec/legacy"
	servicespec "github.com/the-anna-project/spec/service"
	"github.com/xh3b4sd/anna/client/interface/text"
)

// TODO text interface should be a service inside the service collection
func newTextInterface(newServiceCollection servicespec.ServiceCollection, gRPCAddr string) systemspec.TextInterfaceClient {
	newClient := text.New()
	newClient.SetGRPCAddress(gRPCAddr)
	newClient.SetServiceCollection(newServiceCollection)

	return newClient
}
