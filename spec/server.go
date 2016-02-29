package spec

// Server represents a simple bootable object being able to serve network
// resources.
type Server interface {
	// Boot initializes and starts the whole server like booting a machine.
	// Listening to a socket should be done here internally. The call to Boot
	// blocks forever.
	Boot()

	Object
}
