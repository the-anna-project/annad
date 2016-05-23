package spec

// IDType represents some kind of configuration for ID creation.
type IDType int

// IDFactory creates pseudo random hash generation used for ID assignment.
type IDFactory interface {
	// WithType tries to create a new object ID using the given ID type. The
	// returned error might be caused by timeouts reached during the ID creation.
	WithType(idType IDType) (ObjectID, error)
}
