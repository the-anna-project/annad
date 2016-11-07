package spec

// IDType represents some kind of configuration for ID creation.
// TODO remove type when metadata is introduced
type IDType int

// ID creates pseudo random hash generation used for ID assignment.
type ID interface {
	// New tries to create a new object ID using the configured ID type. The
	// returned error might be caused by timeouts reached during the ID creation.
	New() (string, error)

	// WithType tries to create a new object ID using the given ID type. The
	// returned error might be caused by timeouts reached during the ID creation.
	WithType(idType IDType) (string, error)
}
