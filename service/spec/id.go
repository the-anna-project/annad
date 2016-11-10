package spec

// ID creates pseudo random hash generation used for ID assignment.
type ID interface {
	Configure() error

	Metadata() map[string]string

	// New tries to create a new object ID using the configured ID type. The
	// returned error might be caused by timeouts reached during the ID creation.
	New() (string, error)

	Service() Collection

	SetServiceCollection(sc Collection)

	// WithType tries to create a new object ID using the given ID type. The
	// returned error might be caused by timeouts reached during the ID creation.
	WithType(idType int) (string, error)

	Validate() error
}
