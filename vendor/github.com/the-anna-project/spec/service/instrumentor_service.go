package service

import (
	"net/http"

	objectspec "github.com/the-anna-project/spec/object"
)

// InstrumentorService represents an abstraction of instrumentation libraries to
// manage application metrics.
type InstrumentorService interface {
	Boot()
	// ExecFunc wraps basic instrumentation around the given action and executes
	// it.
	//
	// The wrapped action causes the following metric's to be emitted. <prefix>
	// is described by the configured prefix of the current instrumentor.
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
	GetCounter(key string) (objectspec.InstrumentorCounter, error)
	// GetGauge provides a Gauge for the given key. In case there does no
	// counter exist for the given key, one is created.
	GetGauge(key string) (objectspec.InstrumentorGauge, error)
	// GetGauge provides a Gauge for the given key. In case there does no
	// counter exist for the given key, one is created.
	GetHistogram(key string) (objectspec.InstrumentorHistogram, error)
	// GetHTTPEndpoint returns the instrumentor's metric endpoint supposed to
	// be registered in the HTTP server using the instrumentor's HTTP handler.
	GetHTTPEndpoint() string
	// GetHTTPHandler returns the instrumentor's HTTP handler supposed to be
	// registered in the HTTP server using the instrumentor's HTTP endpoint.
	GetHTTPHandler() http.Handler
	// GetPrefixes returns the instrumentor's configured prefix.
	GetPrefixes() []string
	Metadata() map[string]string
	// NewKey returns a new metrics key having all configured prefixes and all
	// given parts properly joined. This could happen e.g. using underscores. In
	// this case it would look as follows.
	//
	//     prefix_prefix_s_s_s_s
	//
	NewKey(s ...string) string
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
	// WrapFunc wraps basic instrumentation around the given action. The returned
	// function can be used as e.g. retry action.
	//
	// The wrapped action causes the same metric's to be emitted as WrapFunc.
	WrapFunc(key string, action func() error) func() error
}
