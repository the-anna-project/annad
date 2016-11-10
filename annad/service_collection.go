package main

import (
	"os"

	"github.com/cenk/backoff"
	kitlog "github.com/go-kit/kit/log"

	"github.com/xh3b4sd/anna/service"
	"github.com/xh3b4sd/anna/service/activator"
	"github.com/xh3b4sd/anna/service/forwarder"
	"github.com/xh3b4sd/anna/service/fs/mem"
	"github.com/xh3b4sd/anna/service/id"
	"github.com/xh3b4sd/anna/service/log"
	"github.com/xh3b4sd/anna/service/network"
	"github.com/xh3b4sd/anna/service/permutation"
	"github.com/xh3b4sd/anna/service/random"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	"github.com/xh3b4sd/anna/service/textinput"
	"github.com/xh3b4sd/anna/service/textoutput"
	"github.com/xh3b4sd/anna/service/tracker"
	systemspec "github.com/xh3b4sd/anna/spec"
)

func newServiceCollection() servicespec.Collection {
	var err error

	// New.
	collection := service.NewCollection()

	// TODO add other services
	activatorService := newActivatorService()
	forwarderService := newForwarderService()
	fileSystemService := newFileSystemService()
	idService := newIDService()
	logService := newLogService()
	networkService := newNetworkService()
	permutationService := newPermutationService()
	randomService := newRandomService()
	textInputService := newTextInputService()
	textOutputService := newTextOutputService()
	trackerService := newTrackerService()

	// Set.
	collection.SetActivator(activatorService)
	collection.SetForwarder(forwarderService)
	collection.SetFS(fileSystemService)
	collection.SetID(idService)
	collection.SetLog(logService)
	collection.SetNetwork(networkService)
	collection.SetPermutation(permutationService)
	collection.SetRandom(randomService)
	collection.SetTextInput(textInputService)
	collection.SetTextOutput(textOutputService)
	collection.SetTracker(trackerService)

	activatorService.SetCollection(collection)
	forwarderService.SetCollection(collection)
	fileSystemService.SetCollection(collection)
	idService.SetCollection(collection)
	logService.SetCollection(collection)
	networkService.SetCollection(collection)
	permutationService.SetCollection(collection)
	randomService.SetCollection(collection)
	textInputService.SetCollection(collection)
	textOutputService.SetCollection(collection)
	trackerService.SetCollection(collection)

	// Validate.
	err = collection.Validate()
	panicOnError(err)

	err = activatorService.Validate()
	panicOnError(err)
	err = forwarderService.Validate()
	panicOnError(err)
	err = fileSystemService.Validate()
	panicOnError(err)
	err = idService.Validate()
	panicOnError(err)
	err = logService.Validate()
	panicOnError(err)
	err = networkService.Validate()
	panicOnError(err)
	err = permutationService.Validate()
	panicOnError(err)
	err = randomService.Validate()
	panicOnError(err)
	err = textInputService.Validate()
	panicOnError(err)
	err = textOutputService.Validate()
	panicOnError(err)
	err = trackerService.Validate()
	panicOnError(err)

	// Configure.
	err = collection.Configure()
	panicOnError(err)

	err = activatorService.Configure()
	panicOnError(err)
	err = forwarderService.Configure()
	panicOnError(err)
	err = fileSystemService.Configure()
	panicOnError(err)
	err = idService.Configure()
	panicOnError(err)
	err = logService.Configure()
	panicOnError(err)
	err = networkService.Configure()
	panicOnError(err)
	err = permutationService.Configure()
	panicOnError(err)
	err = randomService.Configure()
	panicOnError(err)
	err = textInputService.Configure()
	panicOnError(err)
	err = textOutputService.Configure()
	panicOnError(err)
	err = trackerService.Configure()
	panicOnError(err)

	return collection
}

func newActivatorService() servicespec.Activator {
	return activator.New()
}

func newForwarderService() servicespec.Forwarder {
	return forwarder.New()
}

// TODO make mem/os configurable
func newFileSystemService() servicespec.FS {
	return mem.New()
}

func newIDService() servicespec.ID {
	return id.New()
}

func newLogService() servicespec.Log {
	newService := log.New()
	newService.SetRootLogger(kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr)))

	return newService
}

func newNetworkService() servicespec.Network {
	return network.New()
}

func newPermutationService() servicespec.Permutation {
	return permutation.New()
}

func newRandomService() servicespec.Random {
	newService := random.New()

	newService.SetBackoffFactory(func() systemspec.Backoff {
		return backoff.NewExponentialBackOff()
	})

	return newService
}

func newTextInputService() servicespec.TextInput {
	return textinput.New()
}

func newTextOutputService() servicespec.TextOutput {
	return textoutput.New()
}

func newTrackerService() servicespec.Tracker {
	return tracker.New()
}
