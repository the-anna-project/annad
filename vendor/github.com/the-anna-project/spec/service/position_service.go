package service

// PositionService implemnents a service to manage position peers within the
// connection space.
type PositionService interface {
	Boot()
	Default() (string, error)
	Metadata() map[string]string
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
}
