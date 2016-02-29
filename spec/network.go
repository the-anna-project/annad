package spec

// Network implementations provide some sort of higher level business logic.
// They are also used to group and manage neurons. Networks here should not be
// too much aligned with the biological equivalent. Recent implementations
// turned out to be way too complex and stood in the way of getting things
// done. Anyway, networks need to provide critical business logic though. At
// the end the question is what makes sense and what works out.
type Network interface {
	Object

	// Shutdown ends all processes of the network like shutting down a machine.
	// The call to Shutdown blocks until the network is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()

	// Trigger represents the entrance and exit provided for an impulse to walk
	// through the network. Within the network, the impulse might be manipulated.
	Trigger(imp Impulse) (Impulse, error)
}
