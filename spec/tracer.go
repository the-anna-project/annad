package spec

type TraceID string

type Tracer interface {
	GetTraceID() TraceID
}
