package spec

// Log represents a log service used to print log messages.
type Log interface {
	// Line logs a message based on the provided key-value pairs.
	Line(v ...interface{})

	// GetMetadata returns the service's metadata.
	GetMetadata() map[string]string
}

type RootLogger interface {
	Log(v ...interface{}) error
}
