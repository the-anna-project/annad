package main

import (
	"github.com/xh3b4sd/anna/network/tracker"
	"github.com/xh3b4sd/anna/spec"
)

func newTracker(newLog spec.Log, newServiceCollection spec.ServiceCollection, newStorageCollection spec.StorageCollection) (spec.Tracker, error) {
	newTrackerConfig := tracker.DefaultConfig()
	newTrackerConfig.ServiceCollection = newServiceCollection
	newTrackerConfig.Log = newLog
	newTrackerConfig.StorageCollection = newStorageCollection
	newTracker, err := tracker.New(newTrackerConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newTracker, nil
}
