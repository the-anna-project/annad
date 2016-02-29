package spec

// Factory creates new objects on demand to centralize this responsibility.
type Factory interface {
	// Boot initializes and starts the whole factory like booting a machine. The
	// call to Boot blocks until the factory is completely initialized, so you
	// might want to call it in a separate goroutine.
	Boot()

	// NewCore creates a new configured Core object.
	NewCore() (Core, error)

	// NewImpulse creates a new configured Impulse object.
	NewImpulse() (Impulse, error)

	// NewRedisStorage creates a new configured RedisStorage object.
	NewRedisStorage() (Storage, error)

	// NewStrategyNetwork creates a new configured StrategyNetwork object.
	NewStrategyNetwork() (Network, error)

	// Shutdown ends all processes of the factory like shutting down a machine.
	// The call to Shutdown blocks until the factory is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()
}
