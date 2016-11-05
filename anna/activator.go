package main

import (
	"github.com/xh3b4sd/anna/network/activator"
	"github.com/xh3b4sd/anna/spec"
)

func newActivator(newLog spec.Log, newServiceCollection spec.ServiceCollection, newStorageCollection spec.StorageCollection) (spec.Activator, error) {
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
