package spec

// ID creates pseudo random hash generation used for ID assignment.
type ID interface {
	// GetMetadata returns the service's metadata.
	GetMetadata() map[string]string

	// New tries to create a new object ID using the configured ID type. The
	// returned error might be caused by timeouts reached during the ID creation.
	New() (string, error)

	// WithType tries to create a new object ID using the given ID type. The
	// returned error might be caused by timeouts reached during the ID creation.
	WithType(idType int) (string, error)
}
