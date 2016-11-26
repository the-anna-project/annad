package boot

import (
	"os"

	"github.com/cenk/backoff"
	"github.com/garyburd/redigo/redis"
	kitlog "github.com/go-kit/kit/log"

	"github.com/the-anna-project/annad/service/activator"
	"github.com/the-anna-project/annad/service/feature"
	"github.com/the-anna-project/annad/service/forwarder"
	"github.com/the-anna-project/annad/service/network"
	"github.com/the-anna-project/annad/service/tracker"
	servicecollection "github.com/the-anna-project/collection"
	connectionservice "github.com/the-anna-project/connection/service"
	memoryfs "github.com/the-anna-project/fs/memory"
	"github.com/the-anna-project/id"
	inputcollection "github.com/the-anna-project/input/collection"
	textinputservice "github.com/the-anna-project/input/service/text"
	"github.com/the-anna-project/instrumentor/prometheus"
	"github.com/the-anna-project/log"
	outputcollection "github.com/the-anna-project/output/collection"
	textoutputservice "github.com/the-anna-project/output/service/text"
	"github.com/the-anna-project/permutation/service"
	"github.com/the-anna-project/random"
	endpointcollection "github.com/the-anna-project/server/collection"
	metricendpoint "github.com/the-anna-project/server/service/metric"
	textendpoint "github.com/the-anna-project/server/service/text"
	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
	storagecollection "github.com/the-anna-project/storage/collection"
	memorystorage "github.com/the-anna-project/storage/service/memory"
	redisstorage "github.com/the-anna-project/storage/service/redis"
)

func (c *Command) newServiceCollection() servicespec.ServiceCollection {
	// Set.
	collection := servicecollection.New()

	collection.SetActivatorService(c.newActivatorService())
	collection.SetConnectionService(c.newConnectionService())
	collection.SetEndpointCollection(c.newEndpointCollection())
	collection.SetFeatureService(c.newFeatureService())
	collection.SetForwarderService(c.newForwarderService())
	collection.SetFSService(c.newFSService())
	collection.SetIDService(c.newIDService())
	collection.SetInputCollection(c.newInputCollection())
	collection.SetInstrumentorService(c.newInstrumentorService())
	collection.SetLogService(c.newLogService())
	collection.SetNetworkService(c.newNetworkService())
	collection.SetOutputCollection(c.newOutputCollection())
	collection.SetPermutationService(c.newPermutationService())
	collection.SetRandomService(c.newRandomService())
	collection.SetStorageCollection(c.newStorageCollection())
	collection.SetTrackerService(c.newTrackerService())

	collection.Activator().SetServiceCollection(collection)
	collection.Connection().SetServiceCollection(collection)
	collection.Endpoint().Metric().SetServiceCollection(collection)
	collection.Endpoint().Text().SetServiceCollection(collection)
	collection.Feature().SetServiceCollection(collection)
	collection.Forwarder().SetServiceCollection(collection)
	collection.FS().SetServiceCollection(collection)
	collection.ID().SetServiceCollection(collection)
	collection.Input().Text().SetServiceCollection(collection)
	collection.Instrumentor().SetServiceCollection(collection)
	collection.Log().SetServiceCollection(collection)
	collection.Network().SetServiceCollection(collection)
	collection.Output().Text().SetServiceCollection(collection)
	collection.Permutation().SetServiceCollection(collection)
	collection.Random().SetServiceCollection(collection)
	collection.Storage().Connection().SetServiceCollection(collection)
	collection.Storage().Feature().SetServiceCollection(collection)
	collection.Storage().General().SetServiceCollection(collection)
	collection.Tracker().SetServiceCollection(collection)

	return collection
}

func (c *Command) newActivatorService() servicespec.ActivatorService {
	return activator.New()
}

func (c *Command) newConnectionService() servicespec.ConnectionService {
	connectionService := connectionservice.New()

	connectionService.SetDimensionCount(c.configCollection.Space().Dimension().Count())
	connectionService.SetDimensionDepth(c.configCollection.Space().Dimension().Depth())
	connectionService.SetWeight(c.configCollection.Space().Connection().Weight())

	return connectionService
}

func (c *Command) newBackoffFactory() func() objectspec.Backoff {
	return func() objectspec.Backoff {
		return backoff.NewExponentialBackOff()
	}
}

func (c *Command) newEndpointCollection() servicespec.EndpointCollection {
	newCollection := endpointcollection.New()

	metricService := metricendpoint.New()
	metricService.SetAddress(c.configCollection.Endpoint().Metric().Address())

	textService := textendpoint.New()
	textService.SetAddress(c.configCollection.Endpoint().Text().Address())

	newCollection.SetMetricService(metricService)
	newCollection.SetTextService(textService)

	return newCollection
}

func (c *Command) newFeatureService() servicespec.FeatureService {
	return feature.New()
}

func (c *Command) newForwarderService() servicespec.ForwarderService {
	return forwarder.New()
}

// TODO make mem/os configurable
func (c *Command) newFSService() servicespec.FSService {
	return memoryfs.New()
}

func (c *Command) newIDService() servicespec.IDService {
	return id.New()
}

func (c *Command) newInputCollection() servicespec.InputCollection {
	newCollection := inputcollection.New()

	newCollection.SetTextService(textinputservice.New())

	return newCollection
}

func (c *Command) newInstrumentorService() servicespec.InstrumentorService {
	return prometheus.New()
}

func (c *Command) newLayerCollection() servicespec.LayerCollection {
	layerCollection := layercollection.New()

	behaviourService := layerservice.New()
	behaviourService.SetKind("behaviour")
	informationService := layerservice.New()
	informationService.SetKind("information")
	positionService := layerservice.New()
	positionService.SetKind("position")

	layerCollection.SetBehaviourService(behaviourService)
	layerCollection.SetInformationService(informationService)
	layerCollection.SetPositionService(positionService)

	return layerCollection
}

func (c *Command) newLogService() servicespec.LogService {
	logService := log.New()

	logService.SetRootLogger(kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr)))

	return logService
}

func (c *Command) newNetworkService() servicespec.NetworkService {
	return network.New()
}

func (c *Command) newOutputCollection() servicespec.OutputCollection {
	newCollection := outputcollection.New()

	newCollection.SetTextService(textoutputservice.New())

	return newCollection
}

func (c *Command) newPermutationService() servicespec.PermutationService {
	return permutation.New()
}

func (c *Command) newRandomService() servicespec.RandomService {
	randomService := random.New()

	randomService.SetBackoffFactory(c.newBackoffFactory())

	return randomService
}

func (c *Command) newStorageCollection() servicespec.StorageCollection {
	newCollection := storagecollection.New()

	newPool := func(addr string) *redis.Pool {
		newDialConfig := redisstorage.DefaultDialConfig()
		newDialConfig.Addr = addr
		newPoolConfig := redisstorage.DefaultPoolConfig()
		newPoolConfig.Dial = redisstorage.NewDial(newDialConfig)
		newPool := redisstorage.NewPool(newPoolConfig)

		return newPool
	}

	// Connection.
	switch c.configCollection.Storage().Connection().Kind() {
	case "redis":
		connectionService := redisstorage.New()
		connectionService.SetBackoffFactory(c.newBackoffFactory())
		connectionService.SetPool(newPool(c.configCollection.Storage().Connection().Address()))
		connectionService.SetPrefix(c.configCollection.Storage().Connection().Prefix())
		newCollection.SetConnectionService(connectionService)
	case "memory":
		newCollection.SetConnectionService(memorystorage.New())
	default:
		panic(maskAnyf(invalidStorageKindError, "%s", c.configCollection.Storage().Connection().Kind()))
	}

	// Feature.
	switch c.configCollection.Storage().Feature().Kind() {
	case "redis":
		featureService := redisstorage.New()
		featureService.SetBackoffFactory(c.newBackoffFactory())
		featureService.SetPool(newPool(c.configCollection.Storage().Feature().Address()))
		featureService.SetPrefix(c.configCollection.Storage().Feature().Prefix())
		newCollection.SetFeatureService(featureService)
	case "memory":
		newCollection.SetFeatureService(memorystorage.New())
	default:
		panic(maskAnyf(invalidStorageKindError, "%s", c.configCollection.Storage().Feature().Kind()))
	}

	// General.
	switch c.configCollection.Storage().General().Kind() {
	case "redis":
		generalService := redisstorage.New()
		generalService.SetBackoffFactory(c.newBackoffFactory())
		generalService.SetPool(newPool(c.configCollection.Storage().General().Address()))
		generalService.SetPrefix(c.configCollection.Storage().General().Prefix())
		newCollection.SetGeneralService(generalService)
	case "memory":
		newCollection.SetGeneralService(memorystorage.New())
	default:
		panic(maskAnyf(invalidStorageKindError, "%s", c.configCollection.Storage().General().Kind()))
	}

	return newCollection
}

func (c *Command) newTrackerService() servicespec.TrackerService {
	return tracker.New()
}
