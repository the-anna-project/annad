package spec

// Storage represents a persistency management object. Different storages may be
// provided using Provider. Within a receiver function the
// usage of the feature storage may look like this.
//
//     func (n *network) Foo() error {
//       rk, err := n.Storage().Feature().GetRandom()
//       ...
//     }
//
type Storage interface {
	List

	ScoredSet

	Set

	// Shutdown ends all processes of the storage like shutting down a machine.
	// The call to Shutdown blocks until the storage is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()

	StringMap

	String
}
