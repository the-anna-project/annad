package spec

// IDService creates pseudo random hash generation used for ID assignment.
type IDService interface {
	Boot()
	Metadata() map[string]string
	// New tries to create a new object ID using the configured ID type. The
	// returned error might be caused by timeouts reached during the ID creation.
	New() (string, error)
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
	// WithType tries to create a new object ID using the given ID type. The
	// returned error might be caused by timeouts reached during the ID creation.
	WithType(idType int) (string, error)
}
