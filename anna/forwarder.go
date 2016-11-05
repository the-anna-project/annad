package main

import (
	"github.com/xh3b4sd/anna/network/forwarder"
	"github.com/xh3b4sd/anna/spec"
)

func newForwarder(newLog spec.Log, newServiceCollection spec.ServiceCollection, newStorageCollection spec.StorageCollection) (spec.Forwarder, error) {
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
