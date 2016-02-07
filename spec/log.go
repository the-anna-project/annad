package spec

type Tags struct {
	L string

	// O represents the object emitting the log message.
	O Object

	T Tracer

	// V is the verbosity used to log messages.
	V int
}

type Log interface {
	WithTags(tags Tags, f string, v ...interface{})
}
