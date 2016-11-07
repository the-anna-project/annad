package main

import (
	"github.com/xh3b4sd/anna/network/tracker"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

func newTracker(newLog systemspec.Log, newServiceCollection servicespec.Collection, newStorageCollection storagespec.Collection) (systemspec.Tracker, error) {
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
