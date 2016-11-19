package spec

// ServiceCollection represents a collection of factories. This scopes different
// service implementations in a simple container, which can easily be passed
// around.
type ServiceCollection interface {
	Activator() ActivatorService
	Boot()
	Connection() ConnectionService
	Endpoint() EndpointCollection
	Feature() FeatureService
	Forwarder() ForwarderService
	// FSService returns a file system service. It is used to operate on file
	// system abstractions of a certain type.
	FS() FSService
	// IDService returns an ID service. It is used to create IDs of a certain
	// type.
	ID() IDService
	Input() InputCollection
	Instrumentor() InstrumentorService
	// Log returns a log service. It is used to print log messages.
	Log() LogService
	Network() NetworkService
	Output() OutputCollection
	// Permutation returns a permutation service. It is used to permute instances
	// of type PermutationList.
	Permutation() PermutationService
	// Random returns a random service. It is used to create random numbers.
	Random() RandomService
	SetActivatorService(activatorService ActivatorService)
	SetConnectionService(connectionService ConnectionService)
	SetEndpointCollection(endpointCollection EndpointCollection)
	SetFeatureService(featureService FeatureService)
	SetForwarderService(forwarderService ForwarderService)
	SetFSService(fsService FSService)
	SetIDService(idService IDService)
	SetInputCollection(inputCollection InputCollection)
	SetInstrumentorService(instrumentorService InstrumentorService)
	SetLogService(logService LogService)
	SetNetworkService(networkService NetworkService)
	SetOutputCollection(outputCollection OutputCollection)
	SetPermutationService(permutationService PermutationService)
	SetRandomService(randomService RandomService)
	SetStorageCollection(storageCollection StorageCollection)
	SetTrackerService(trackerService TrackerService)
	// Shutdown ends all processes of the service collection like shutting down a
	// machine. The call to Shutdown blocks until the service collection is
	// completely shut down, so you might want to call it in a separate goroutine.
	Shutdown()
	Storage() StorageCollection
	Tracker() TrackerService
}
