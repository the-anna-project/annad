package spec

// CLGIndex represents the CLG index. It provides a CLG profile generator used
// to generate CLG profiles based on the CLGs provided by the CLG collection.
type CLGIndex interface {
	// Boot initializes and starts the whole CLG index like booting a machine.
	// The call to Boot blocks until the CLG index is completely initialized, so
	// you might want to call it in a separate goroutine.
	Boot()

	// CreateProfiles uses the given CLG profile generator to create CLG profiles
	// obtained by the CLG collection, whether there are already proper profiles.
	// In case a proper profile exists, it stays untouched. In case a profile is
	// outdated or not complete, it will be modified. A CLG profile is used to
	// describe a CLG's state, e.g. its interface, functionality and identity.
	CreateProfiles(generator CLGProfileGenerator) error

	// GetGenerator returns the CLG profile generator configured for the current
	// CLG index.
	GetGenerator() CLGProfileGenerator

	Object

	// Shutdown ends all processes of the CLG index like shutting down a machine.
	// The call to Shutdown blocks until the CLG index is completely shut down,
	// so you might want to call it in a separate goroutine.
	Shutdown()
}
