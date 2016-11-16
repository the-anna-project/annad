package spec

// LogService represents a log service used to print log messages.
type LogService interface {
	Boot()
	// Line logs a message based on the provided key-value pairs.
	Line(v ...interface{})
	Metadata() map[string]string
	Service() ServiceCollection
	SetRootLogger(rootLogger RootLogger)
	SetServiceCollection(serviceCollection ServiceCollection)
}

// RootLogger is the underlying logger actually printing log messages.
type RootLogger interface {
	Log(v ...interface{}) error
}
