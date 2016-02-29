package spec

// TraceID represents an identifier that is comparable and comprehensible
// across requests.
type TraceID string

// Tracer represents a container for comparable and comprehensible data. A
// tracer can be used in middlewares to differenciate or map requests.
type Tracer interface {
	// GetTraceID returns the tracer's ID.
	GetTraceID() TraceID
}
