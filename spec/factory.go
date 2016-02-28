package spec

type Factory interface {
	// Boot initializes and starts the whole factory like booting a machine. The
	// call to Boot blocks until the factory is completely initialized, so you
	// might want to call it in a separate goroutine.
	Boot()

	NewCore() (Core, error)

	NewImpulse() (Impulse, error)

	NewRedisStorage() (Storage, error)

	NewStrategyNetwork() (Network, error)

	// Shutdown ends all processes of the factory like shutting down a machine.
	// The call to Shutdown blocks until the factory is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()
}
