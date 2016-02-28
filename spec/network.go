package spec

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
