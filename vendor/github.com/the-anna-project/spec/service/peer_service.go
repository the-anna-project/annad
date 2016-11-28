package service

// PeerService implemnents a service to manage peers within the connection
// space.
type PeerService interface {
	Boot()
	Create(peer string) error
	Delete(peer string) error
	Metadata() map[string]string
	Search(peer string) (map[string]string, error)
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
}
