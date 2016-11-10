package main

import (
	"github.com/xh3b4sd/anna/client/interface/text"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
)

// TODO text interface should be a service inside the service collection
func newTextInterface(newServiceCollection servicespec.Collection, gRPCAddr string) systemspec.TextInterfaceClient {
	newClient := text.New()
	newClient.SetGRPCAddress(gRPCAddr)
	newClient.SetServiceCollection(newServiceCollection)

	err := newClient.Validate()
	panicOnError(err)

	err = newClient.Configure()
	panicOnError(err)

	return newClient
}
