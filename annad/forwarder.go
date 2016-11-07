package main

import (
	"github.com/xh3b4sd/anna/network/forwarder"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

func newForwarder(newLog systemspec.Log, newServiceCollection servicespec.Collection, newStorageCollection storagespec.Collection) (systemspec.Forwarder, error) {
	newForwarderConfig := forwarder.DefaultConfig()
	newForwarderConfig.ServiceCollection = newServiceCollection
	newForwarderConfig.Log = newLog
	newForwarderConfig.StorageCollection = newStorageCollection
	newForwarder, err := forwarder.New(newForwarderConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newForwarder, nil
}
