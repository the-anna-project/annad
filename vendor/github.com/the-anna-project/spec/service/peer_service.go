package service

// PeerService implemnents a service to manage peers within the connection
// space.
//
// Following is a list of properties each peer has applied in form of metadata
// to itself.
//
//     created
//
type PeerService interface {
	Boot()
	// Create creates a new peer for the given peer value.
	Create(peer string) error
	// Delete deletes the peer identified by the given peer value.
	Delete(peer string) error
	Metadata() map[string]string
	// Search returns all metadata associated with the peer identified by the
	// given peer value.
	Search(peer string) (map[string]string, error)
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
	Shutdown()
}
