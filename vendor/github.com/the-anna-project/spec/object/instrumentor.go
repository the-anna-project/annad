package spec

// InstrumentorCounter is a metric that can be arbitrarily incremented.
type InstrumentorCounter interface {
	// IncrBy increments the current counter by the given delta.
	IncrBy(delta float64)
}

// InstrumentorGauge is a metric that can be arbitrarily incremented or
// decremented.
type InstrumentorGauge interface {
	// DecrBy decrements the current gauge by the given delta.
	DecrBy(delta float64)
	// IncrBy increments the current gauge by the given delta.
	IncrBy(delta float64)
}

// InstrumentorHistogram is a metric to observe samples over time.
type InstrumentorHistogram interface {
	// Observe tracks the given sample used for aggregation of the current
	// histogramm.
	Observe(sample float64)
}
