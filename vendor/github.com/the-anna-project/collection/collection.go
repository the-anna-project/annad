package collection

import (
	"sync"

	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new service collection.
func New() servicespec.ServiceCollection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	activatorService    servicespec.ActivatorService
	connectionService   servicespec.ConnectionService
	endpointCollection  servicespec.EndpointCollection
	featureService      servicespec.FeatureService
	forwarderService    servicespec.ForwarderService
	fsService           servicespec.FSService
	idService           servicespec.IDService
	inputCollection     servicespec.InputCollection
	instrumentorService servicespec.InstrumentorService
	layerCollection     servicespec.LayerCollection
	logService          servicespec.LogService
	networkService      servicespec.NetworkService
	outputCollection    servicespec.OutputCollection
	permutationService  servicespec.PermutationService
	randomService       servicespec.RandomService
	storageCollection   servicespec.StorageCollection
	trackerService      servicespec.TrackerService
	workerService       servicespec.WorkerService

	// Settings.

	shutdownOnce sync.Once
}

func (c *collection) Activator() servicespec.ActivatorService {
	return c.activatorService
}

func (c *collection) Boot() {
	go c.Activator().Boot()
	go c.Connection().Boot()
	go c.Endpoint().Boot()
	go c.Feature().Boot()
	go c.Forwarder().Boot()
	go c.FS().Boot()
	go c.ID().Boot()
	go c.Input().Boot()
	go c.Instrumentor().Boot()
	go c.Log().Boot()
	go c.Network().Boot()
	go c.Output().Boot()
	go c.Permutation().Boot()
	go c.Random().Boot()
	go c.Storage().Boot()
	go c.Tracker().Boot()
	go c.Worker().Boot()
}

func (c *collection) Connection() servicespec.ConnectionService {
	return c.connectionService
}

func (c *collection) Endpoint() servicespec.EndpointCollection {
	return c.endpointCollection
}

func (c *collection) Feature() servicespec.FeatureService {
	return c.featureService
}

func (c *collection) Forwarder() servicespec.ForwarderService {
	return c.forwarderService
}

func (c *collection) FS() servicespec.FSService {
	return c.fsService
}

func (c *collection) ID() servicespec.IDService {
	return c.idService
}

func (c *collection) Input() servicespec.InputCollection {
	return c.inputCollection
}

func (c *collection) Instrumentor() servicespec.InstrumentorService {
	return c.instrumentorService
}

func (c *collection) Layer() servicespec.LayerCollection {
	return c.layerCollection
}

func (c *collection) Log() servicespec.LogService {
	return c.logService
}

func (c *collection) Network() servicespec.NetworkService {
	return c.networkService
}

func (c *collection) Output() servicespec.OutputCollection {
	return c.outputCollection
}

func (c *collection) Permutation() servicespec.PermutationService {
	return c.permutationService
}

func (c *collection) Random() servicespec.RandomService {
	return c.randomService
}

func (c *collection) SetActivatorService(activator servicespec.ActivatorService) {
	c.activatorService = activator
}

func (c *collection) SetConnectionService(connectionService servicespec.ConnectionService) {
	c.connectionService = connectionService
}

func (c *collection) SetEndpointCollection(endpointCollection servicespec.EndpointCollection) {
	c.endpointCollection = endpointCollection
}

func (c *collection) SetFeatureService(featureService servicespec.FeatureService) {
	c.featureService = featureService
}

func (c *collection) SetForwarderService(forwarderService servicespec.ForwarderService) {
	c.forwarderService = forwarderService
}

func (c *collection) SetFSService(fsService servicespec.FSService) {
	c.fsService = fsService
}

func (c *collection) SetIDService(idService servicespec.IDService) {
	c.idService = idService
}

func (c *collection) SetInputCollection(inputCollection servicespec.InputCollection) {
	c.inputCollection = inputCollection
}

func (c *collection) SetInstrumentorService(instrumentorService servicespec.InstrumentorService) {
	c.instrumentorService = instrumentorService
}

func (c *collection) SetLayerCollection(layerCollection servicespec.LayerCollection) {
	c.layerCollection = layerCollection
}

func (c *collection) SetLogService(logService servicespec.LogService) {
	c.logService = logService
}

func (c *collection) SetNetworkService(networkService servicespec.NetworkService) {
	c.networkService = networkService
}

func (c *collection) SetOutputCollection(outputCollection servicespec.OutputCollection) {
	c.outputCollection = outputCollection
}

func (c *collection) SetPermutationService(permutationService servicespec.PermutationService) {
	c.permutationService = permutationService
}

func (c *collection) SetRandomService(randomService servicespec.RandomService) {
	c.randomService = randomService
}

func (c *collection) SetStorageCollection(storageCollection servicespec.StorageCollection) {
	c.storageCollection = storageCollection
}

func (c *collection) SetTrackerService(trackerService servicespec.TrackerService) {
	c.trackerService = trackerService
}

func (c *collection) SetWorkerService(workerService servicespec.WorkerService) {
	c.workerService = workerService
}

func (c *collection) Shutdown() {
	c.shutdownOnce.Do(func() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			c.Endpoint().Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			c.Network().Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			c.Storage().Shutdown()
			wg.Done()
		}()

		wg.Wait()
	})
}

func (c *collection) Storage() servicespec.StorageCollection {
	return c.storageCollection
}

func (c *collection) Tracker() servicespec.TrackerService {
	return c.trackerService
}

func (c *collection) Worker() servicespec.WorkerService {
	return c.workerService
}
