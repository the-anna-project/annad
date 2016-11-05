package spec

// ServiceCollection represents a collection of factories. This scopes different
// service implementations in a simple container, which can easily be passed
// around.
type ServiceCollection interface {
	// FS returns a file system service. It is used to operate on file system
	// abstractions of a certain type.
	FS() FileSystem

	// ID returns an ID service. It is used to create IDs of a certain type.
	ID() ID

	// Permutation returns a permutation service. It is used to permute instances
	// of type PermutationList.
	Permutation() Permutation

	// Random returns a random service. It is used to create random numbers.
	Random() Random

	// Shutdown ends all processes of the service collection like shutting down a
	// machine. The call to Shutdown blocks until the service collection is
	// completely shut down, so you might want to call it in a separate goroutine.
	Shutdown()
}

// ServiceProvider should be implemented by every object which wants to use
// factories. This then creates an API between service implementations and
// service users.
type ServiceProvider interface {
	Service() ServiceCollection
}
