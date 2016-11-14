package spec

import (
	"net/http"
)

// Instrumentor represents an abstraction of instrumentation libraries to
// manage application metrics.
type Instrumentor interface {
	Configure() error

	// ExecFunc wraps basic instrumentation around the given action and executes
	// it.
	//
	// The wrapped action causes the following metric's to be emitted. <prefix>
	// is described by the configured prefix of the current Instrumentor.
	//
	//     <prefix>_<key>_durations_histogram_milliseconds
	//
	//         Holds the action's duration in milliseconds. This metric is
	//         emitted for each executed action.
	//
	//     <prefix>_<key>_errors_total
	//
	//         Holds the action's error count. This metric is emitted for each
	//         error returned by the given action.
	//
	ExecFunc(key string, action func() error) error

	// GetCounter provides a Counter for the given key. In case there does no
	// counter exist for the given key, one is created.
	GetCounter(key string) (Counter, error)

	// GetGauge provides a Gauge for the given key. In case there does no
	// counter exist for the given key, one is created.
	GetGauge(key string) (Gauge, error)

	// GetGauge provides a Gauge for the given key. In case there does no
	// counter exist for the given key, one is created.
	GetHistogram(key string) (Histogram, error)

	// GetHTTPEndpoint returns the Instrumentor's metric endpoint supposed to
	// be registered in the HTTP server using the Instrumentor's HTTP handler.
	GetHTTPEndpoint() string

	// GetHTTPHandler returns the Instrumentor's HTTP handler supposed to be
	// registered in the HTTP server using the Instrumentor's HTTP endpoint.
	GetHTTPHandler() http.Handler

	// GetPrefixes returns the Instrumentor's configured prefix.
	GetPrefixes() []string

	Metadata() map[string]string

	// NewKey returns a new metrics key having all configured prefixes and all
	// given parts properly joined. This could happen e.g. using underscores. In
	// this case it would look as follows.
	//
	//     prefix_prefix_s_s_s_s
	//
	NewKey(s ...string) string

	Service() Collection

	SetServiceCollection(serviceCollection Collection)



	// WrapFunc wraps basic instrumentation around the given action. The returned
	// function can be used as e.g. retry action.
	//
	// The wrapped action causes the same metric's to be emitted as WrapFunc.
	WrapFunc(key string, action func() error) func() error
}

// Counter is a metric that can be arbitrarily incremented.
type Counter interface {
	// IncrBy increments the current counter by the given delta.
	IncrBy(delta float64)
}

// Gauge is a metric that can be arbitrarily incremented or decremented.
type Gauge interface {
	// DecrBy decrements the current gauge by the given delta.
	DecrBy(delta float64)

	// IncrBy increments the current gauge by the given delta.
	IncrBy(delta float64)
}

// Histogram is a metric to observe samples over time.
type Histogram interface {
	// Observe tracks the given sample used for aggregation of the current
	// histogramm.
	Observe(sample float64)
}
