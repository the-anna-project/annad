package spec

// Log represents a log service used to print log messages.
type Log interface {
	Configure() error

	// Line logs a message based on the provided key-value pairs.
	Line(v ...interface{})

	Metadata() map[string]string

	Service() Collection

	SetRootLogger(rl RootLogger)
	SetServiceCollection(sc Collection)

	Validate() error
}

type RootLogger interface {
	Log(v ...interface{}) error
}
