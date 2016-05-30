package spec

// Server represents a simple bootable object being able to serve network
// resources.
type Server interface {
	// Boot initializes and starts the whole server like booting a machine.
	// Listening to a socket should be done here internally. The call to Boot
	// blocks forever.
	Boot()

	Object

	// Shutdown ends all processes of the server like shutting down a machine.
	// The call to Shutdown blocks until the server is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()
}
