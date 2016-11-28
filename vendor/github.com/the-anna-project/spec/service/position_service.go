package service

// PositionService implemnents a service to manage position peers within the
// connection space.
type PositionService interface {
	Boot()
	// Create creates the position peer for the associated peer. In addition
	// create returns the position peer associated with the provided peer.
	Create(peer string) (string, error)
	// Delete deletes the position peer associated with the given peer. In
	// addition delete returns the position peer associated with the provided
	// peer.
	Delete(peer string) (string, error)
	Metadata() map[string]string
	// Search returns the position peer associated with the provided peer.
	Search(peer string) (string, error)
	Service() ServiceCollection
	SetDimensionCount(dimensionCount int)
	SetDimensionDepth(dimensionDepth int)
	SetServiceCollection(serviceCollection ServiceCollection)
}
