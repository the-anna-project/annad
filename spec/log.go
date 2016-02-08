package spec

// Tags provides criteria to decide what log messages are supposed to be
// logged. Emitted logs not matching the given critera by Tags are not supposed
// to be logged.
type Tags struct {
	// L is the log level. E.g. debug or error.
	L string

	// O represents the object emitting the log message.
	O Object

	// Tracer represents some tracable context passed through. Logs related to a
	// specific trace ID should be caused by a common request.
	T Tracer

	// V is the verbosity used to log messages.
	V int
}

// Log is a logger used to filter logs based on tags before actually logging
// them.
type Log interface {
	WithTags(tags Tags, f string, v ...interface{})
}

// RootLogger is the underlying logger used to actually log messages.
type RootLogger interface {
	Println(v ...interface{})
}
