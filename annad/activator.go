package main

import (
	"github.com/xh3b4sd/anna/network/activator"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

func newActivator(newLog systemspec.Log, newServiceCollection servicespec.Collection, newStorageCollection storagespec.Collection) (systemspec.Activator, error) {
	newActivatorConfig := activator.DefaultConfig()
	newActivatorConfig.ServiceCollection = newServiceCollection
	newActivatorConfig.Log = newLog
	newActivatorConfig.StorageCollection = newStorageCollection
	newActivator, err := activator.New(newActivatorConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newActivator, nil
}
