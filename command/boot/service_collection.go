package boot

import (
	"os"

	"github.com/cenk/backoff"
	kitlog "github.com/go-kit/kit/log"

	"github.com/xh3b4sd/anna/service"
	"github.com/xh3b4sd/anna/service/activator"
	"github.com/xh3b4sd/anna/service/connection"
	"github.com/xh3b4sd/anna/service/endpoint"
	"github.com/xh3b4sd/anna/service/endpoint/metric"
	"github.com/xh3b4sd/anna/service/endpoint/text"
	"github.com/xh3b4sd/anna/service/feature"
	"github.com/xh3b4sd/anna/service/forwarder"
	"github.com/xh3b4sd/anna/service/fs/mem"
	"github.com/xh3b4sd/anna/service/id"
	"github.com/xh3b4sd/anna/service/instrumentor/prometheus"
	"github.com/xh3b4sd/anna/service/log"
	"github.com/xh3b4sd/anna/service/network"
	"github.com/xh3b4sd/anna/service/permutation"
	"github.com/xh3b4sd/anna/service/random"
	"github.com/xh3b4sd/anna/service/spec"
	"github.com/xh3b4sd/anna/service/storage"
	"github.com/xh3b4sd/anna/service/storage/memory"
	"github.com/xh3b4sd/anna/service/storage/redis"
	"github.com/xh3b4sd/anna/service/textinput"
	"github.com/xh3b4sd/anna/service/textoutput"
	"github.com/xh3b4sd/anna/service/tracker"
)

func (c *Command) newServiceCollection() spec.Collection {
	// Set.
	collection := service.NewCollection()

	collection.SetActivator(c.newActivatorService())
	collection.SetConnection(c.newConnectionService())
	collection.SetEndpointCollection(c.newEndpointCollection())
	collection.SetFeature(c.newFeatureService())
	collection.SetForwarder(c.newForwarderService())
	collection.SetFS(c.newFSService())
	collection.SetID(c.newIDService())
	collection.SetInstrumentor(c.newInstrumentorService())
	collection.SetLog(c.newLogService())
	collection.SetNetwork(c.newNetworkService())
	collection.SetPermutation(c.newPermutationService())
	collection.SetRandom(c.newRandomService())
	collection.SetStorageCollection(c.newStorageCollection())
	collection.SetTextInput(c.newTextInputService())
	collection.SetTextOutput(c.newTextOutputService())
	collection.SetTracker(c.newTrackerService())

	collection.Activator().SetServiceCollection(collection)
	collection.Connection().SetServiceCollection(collection)
	collection.Endpoint().Metric().SetServiceCollection(collection)
	collection.Endpoint().Text().SetServiceCollection(collection)
	collection.Feature().SetServiceCollection(collection)
	collection.Forwarder().SetServiceCollection(collection)
	collection.FS().SetServiceCollection(collection)
	collection.ID().SetServiceCollection(collection)
	collection.Instrumentor().SetServiceCollection(collection)
	collection.Log().SetServiceCollection(collection)
	collection.Network().SetServiceCollection(collection)
	collection.Permutation().SetServiceCollection(collection)
	collection.Random().SetServiceCollection(collection)
	collection.Storage().Connection().SetServiceCollection(collection)
	collection.Storage().Feature().SetServiceCollection(collection)
	collection.Storage().General().SetServiceCollection(collection)
	collection.TextInput().SetServiceCollection(collection)
	collection.TextOutput().SetServiceCollection(collection)
	collection.Tracker().SetServiceCollection(collection)

	return collection
}

func (c *Command) newActivatorService() spec.Activator {
	return activator.New()
}

func (c *Command) newConnectionService() spec.Connection {
	newService := connection.New()

	newService.SetDimensionCount(c.configCollection.Space().Dimension().Count())
	newService.SetDimensionDepth(c.configCollection.Space().Dimension().Depth())
	newService.SetWeight(c.configCollection.Space().Connection().Weight())

	return newService
}

func (c *Command) newBackoffFactory() func() spec.Backoff {
	return func() spec.Backoff {
		return backoff.NewExponentialBackOff()
	}
}

func (c *Command) newEndpointCollection() spec.EndpointCollection {
	newCollection := endpoint.NewCollection()

	metricService := metric.New()
	metricService.SetAddress(c.configCollection.Endpoint().Metric().Address())

	textService := text.New()
	textService.SetAddress(c.configCollection.Endpoint().Text().Address())

	newCollection.SetMetric(metricService)
	newCollection.SetText(textService)

	return newCollection
}

func (c *Command) newFeatureService() spec.Feature {
	return feature.New()
}

func (c *Command) newForwarderService() spec.Forwarder {
	return forwarder.New()
}

// TODO make mem/os configurable
func (c *Command) newFSService() spec.FS {
	return mem.New()
}

func (c *Command) newIDService() spec.ID {
	return id.New()
}

func (c *Command) newInstrumentorService() spec.Instrumentor {
	return prometheus.New()
}

func (c *Command) newLogService() spec.Log {
	newService := log.New()

	newService.SetRootLogger(kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr)))

	return newService
}

func (c *Command) newNetworkService() spec.Network {
	return network.New()
}

func (c *Command) newPermutationService() spec.Permutation {
	return permutation.New()
}

func (c *Command) newRandomService() spec.Random {
	newService := random.New()

	newService.SetBackoffFactory(c.newBackoffFactory())

	return newService
}

func (c *Command) newStorageCollection() spec.StorageCollection {
	newCollection := storage.NewCollection()

	// Connection.
	switch c.configCollection.Storage().Connection().Kind() {
	case "redis":
		newService := redis.New()
		newService.SetAddress(c.configCollection.Storage().Connection().Address())
		newService.SetBackoffFactory(c.newBackoffFactory())
		newService.SetPrefix(c.configCollection.Storage().Connection().Prefix())
		newCollection.SetConnection(newService)
	case "memory":
		newCollection.SetConnection(memory.New())
	default:
		panic(maskAnyf(invalidStorageKindError, "%s", c.configCollection.Storage().Connection().Kind()))
	}

	// Feature.
	switch c.configCollection.Storage().Feature().Kind() {
	case "redis":
		newService := redis.New()
		newService.SetAddress(c.configCollection.Storage().Feature().Address())
		newService.SetBackoffFactory(c.newBackoffFactory())
		newService.SetPrefix(c.configCollection.Storage().Feature().Prefix())
		newCollection.SetConnection(newService)
	case "memory":
		newCollection.SetFeature(memory.New())
	default:
		panic(maskAnyf(invalidStorageKindError, "%s", c.configCollection.Storage().Feature().Kind()))
	}

	// General.
	switch c.configCollection.Storage().General().Kind() {
	case "redis":
		newService := redis.New()
		newService.SetAddress(c.configCollection.Storage().General().Address())
		newService.SetBackoffFactory(c.newBackoffFactory())
		newService.SetPrefix(c.configCollection.Storage().General().Prefix())
		newCollection.SetConnection(newService)
	case "memory":
		newCollection.SetGeneral(memory.New())
	default:
		panic(maskAnyf(invalidStorageKindError, "%s", c.configCollection.Storage().General().Kind()))
	}

	return newCollection
}

func (c *Command) newTextInputService() spec.TextInput {
	return textinput.New()
}

func (c *Command) newTextOutputService() spec.TextOutput {
	return textoutput.New()
}

func (c *Command) newTrackerService() spec.Tracker {
	return tracker.New()
}
