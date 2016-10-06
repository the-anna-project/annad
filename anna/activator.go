package main

import (
	"github.com/xh3b4sd/anna/network/activator"
	"github.com/xh3b4sd/anna/spec"
)

func newActivator(newLog spec.Log, newFactoryCollection spec.FactoryCollection, newStorageCollection spec.StorageCollection) (spec.Activator, error) {
	newActivatorConfig := activator.DefaultConfig()
	newActivatorConfig.FactoryCollection = newFactoryCollection
	newActivatorConfig.Log = newLog
	newActivatorConfig.StorageCollection = newStorageCollection
	newActivator, err := activator.New(newActivatorConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newActivator, nil
}
