package spec

// Tags provides criteria to decide what log messages are supposed to be
// logged. Emitted logs not matching the given criteria by Tags are not supposed
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
	// Register adds the given object type to the list of known objects allowed
	// to emit log messages. This information is used to make filtering via
	// object types possible.
	Register(objectType ObjectType) error

	// ResetLevels sets the list of log levels back to its default value.
	ResetLevels() error

	// ResetObjects sets the list of log objects back to its default value.
	ResetObjects() error

	// ResetVerbosity sets the log verbosity back to its default value.
	ResetVerbosity() error

	// SetLevels takes a comma separated list of provided log levels and causes
	// the logger to only log messages tagged related to log levels of the given
	// list.
	SetLevels(list string) error

	// SetObjects takes a comma separated list of provided object types and
	// causes the logger to only log messages tagged related to object types of
	// the given list.
	SetObjects(list string) error

	// SetVerbosity causes the logger to only log messages tagged related to the
	// given verbosity.
	SetVerbosity(verbosity int) error

	// WithTags logs a message based on the provided tags.
	WithTags(tags Tags, f string, v ...interface{})
}

// RootLogger is the underlying logger used to actually log messages.
type RootLogger interface {
	// Println just takes an arbitrary list of arguments and prints a line to the
	// configured output.
	Println(v ...interface{})
}
