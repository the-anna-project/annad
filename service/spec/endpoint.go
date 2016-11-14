package spec

// Endpoint represents a simple bootable object being able to serve network
// resources.
type Endpoint interface {
	// Boot initializes and starts the whole server like booting a machine.
	// Listening to a socket should be done here internally. The call to Boot
	// blocks forever.
	Boot()
	Service() Collection
	SetAddress(address string)
	SetServiceCollection(serviceCollection Collection)
	// Shutdown ends all processes of the server like shutting down a machine.
	// The call to Shutdown blocks until the server is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()
}

// EndpointCollection represents a collection of endpoint instances. This scopes
// different endpoint implementations in a simple container, which can easily be
// passed around.
type EndpointCollection interface {
	Boot()
	Metric() Endpoint
	Text() Endpoint
	Shutdown()
}
