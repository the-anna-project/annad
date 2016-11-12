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
	"github.com/xh3b4sd/anna/service/instrumentor/prometheus"
	"github.com/xh3b4sd/anna/service/log"
	"github.com/xh3b4sd/anna/service/network"
	"github.com/xh3b4sd/anna/service/permutation"
	"github.com/xh3b4sd/anna/service/random"
	"github.com/xh3b4sd/anna/service/server"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	"github.com/xh3b4sd/anna/service/textendpoint"
	"github.com/xh3b4sd/anna/service/textinput"
	"github.com/xh3b4sd/anna/service/textoutput"
	"github.com/xh3b4sd/anna/service/tracker"
)

func (a *annactl) newServiceCollection() servicespec.Collection {
	var err error

	// Set.
	collection := service.NewCollection()

	// TODO add other services

	collection.SetActivator(a.newActivatorService())
	collection.SetForwarder(a.newForwarderService())
	collection.SetFS(a.newFSService())
	collection.SetID(a.newIDService())
	collection.SetID(a.newInstrumentorService())
	collection.SetLog(a.newLogService())
	collection.SetNetwork(a.newNetworkService())
	collection.SetPermutation(a.newPermutationService())
	collection.SetRandom(a.newRandomService())
	collection.SetServer(a.newServerService())
	collection.SetTextEndpoint(a.newTextEndpointService())
	collection.SetTextInput(a.newTextInputService())
	collection.SetTextOutput(a.newTextOutputService())
	collection.SetTracker(a.newTrackerService())

	collection.Activator().SetServiceCollection(collection)
	collection.Forwarder().SetServiceCollection(collection)
	collection.FS().SetServiceCollection(collection)
	collection.ID().SetServiceCollection(collection)
	collection.Instrumentor().SetServiceCollection(collection)
	collection.Log().SetServiceCollection(collection)
	collection.Network().SetServiceCollection(collection)
	collection.Permutation().SetServiceCollection(collection)
	collection.Random().SetServiceCollection(collection)
	collection.Server().SetServiceCollection(collection)
	collection.TextEndpoint().SetServiceCollection(collection)
	collection.TextInput().SetServiceCollection(collection)
	collection.TextOutput().SetServiceCollection(collection)
	collection.Tracker().SetServiceCollection(collection)

	// Validate.
	panicOnError(collection.Validate())

	panicOnError(collection.Activator().Validate())
	panicOnError(collection.Forwarder().Validate())
	panicOnError(collection.FS().Validate())
	panicOnError(collection.ID().Validate())
	panicOnError(collection.Instrumentor().Validate())
	panicOnError(collection.Log().Validate())
	panicOnError(collection.Network().Validate())
	panicOnError(collection.Permutation().Validate())
	panicOnError(collection.Random().Validate())
	panicOnError(collection.Server().Validate())
	panicOnError(collection.TextEndpoint().Validate())
	panicOnError(collection.TextInput().Validate())
	panicOnError(collection.TextOutput().Validate())
	panicOnError(collection.Tracker().Validate())

	// Configure.
	panicOnError(collection.Configure())

	panicOnError(collection.Activator().Configure())
	panicOnError(collection.Forwarder().Configure())
	panicOnError(collection.FS().Configure())
	panicOnError(collection.ID().Configure())
	panicOnError(collection.Instrumentor().Configure())
	panicOnError(collection.Log().Configure())
	panicOnError(collection.Network().Configure())
	panicOnError(collection.Permutation().Configure())
	panicOnError(collection.Random().Configure())
	panicOnError(collection.Server().Configure())
	panicOnError(collection.TextEndpoint().Configure())
	panicOnError(collection.TextInput().Configure())
	panicOnError(collection.TextOutput().Configure())
	panicOnError(collection.Tracker().Configure())

	return collection
}

func (a *annactl) newActivatorService() servicespec.Activator {
	return activator.New()
}

func (a *annactl) newForwarderService() servicespec.Forwarder {
	return forwarder.New()
}

// TODO make mem/os configurable
func (a *annactl) newFSService() servicespec.FS {
	return mem.New()
}

func (a *annactl) newIDService() servicespec.ID {
	return id.New()
}

func (a *annactl) newInstrumentorService() servicespec.Instrumentor {
	return prometheus.New()
}

func (a *annactl) newLogService() servicespec.Log {
	newService := log.New()
	newService.SetRootLogger(kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr)))

	return newService
}

func (a *annactl) newNetworkService() servicespec.Network {
	return network.New()
}

func (a *annactl) newPermutationService() servicespec.Permutation {
	return permutation.New()
}

func (a *annactl) newRandomService() servicespec.Random {
	newService := random.New()

	newService.SetBackoffFactory(func() servicespec.Backoff {
		return backoff.NewExponentialBackOff()
	})

	return newService
}

func (a *annactl) newServerService() servicespec.Server {
	newService := server.New()

	newService.SetHTTPAddress(a.flags.HTTPAddr)

	return newService
}

func (a *annactl) newTextEndpointService() servicespec.TextEndpoint {
	newService := textendpoint.New()

	newService.SetGRPCAddress(a.flags.GRPCAddr)

	return newService
}

func (a *annactl) newTextInputService() servicespec.TextInput {
	return textinput.New()
}

func (a *annactl) newTextOutputService() servicespec.TextOutput {
	return textoutput.New()
}

func (a *annactl) newTrackerService() servicespec.Tracker {
	return tracker.New()
}
