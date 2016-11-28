package service

// PeerService implemnents a service to manage peers within the connection
// space.
type PeerService interface {
	Boot()
	Create(peerA, peerB string) error
	Delete(peerA string) error
	Metadata() map[string]string
	Search(peer string) ([]string, error)
	Service() ServiceCollection
	SetDimensionCount(dimensionCount int)
	SetDimensionDepth(dimensionDepth int)
	SetServiceCollection(serviceCollection ServiceCollection)
}
