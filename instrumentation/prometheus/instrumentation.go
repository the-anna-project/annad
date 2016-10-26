package prometheus

import (
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
)

// InstrumentationConfig represents the configuration used to create a new
// prometheus instrumentation object.
type InstrumentationConfig struct {
	// Settings.

	// Prefixes represents the Instrumentation's ordered prefixes. It is
	// recommended to use the following prefixes.
	//
	//     []string{"anna", "<prefix>"}
	//
	Prefixes []string

	// HTTPEndpoint represents the HTTP endpoint used to register the
	// HTTPHandler. In the context of Prometheus this is usually /metrics.
	HTTPEndpoint string

	// HTTPHandler represents the HTTP handler used to register the Prometheus
	// registry in the HTTP server.
	HTTPHandler http.Handler
}

// DefaultInstrumentationConfig provides a default configuration to create a
// new prometheus instrumentation object by best effort.
func DefaultInstrumentationConfig() InstrumentationConfig {
	newConfig := InstrumentationConfig{
		// Settings.
		Prefixes:     []string{"anna"},
		HTTPEndpoint: "/metrics",
		HTTPHandler:  prometheus.Handler(),
	}

	return newConfig
}

// NewInstrumentation creates a new configured prometheus instrumentation object.
func NewInstrumentation(config InstrumentationConfig) (spec.Instrumentation, error) {
	newPrometheus := &instrumentation{
		InstrumentationConfig: config,

		Counters:   map[string]spec.Counter{},
		Gauges:     map[string]spec.Gauge{},
		Histograms: map[string]spec.Histogram{},

		Mutex: sync.Mutex{},
	}

	if len(newPrometheus.Prefixes) == 0 {
		return nil, maskAnyf(invalidConfigError, "prefixes must not be empty")
	}
	if newPrometheus.HTTPEndpoint == "" {
		return nil, maskAnyf(invalidConfigError, "HTTP endpoint must not be empty")
	}
	if newPrometheus.HTTPHandler == nil {
		return nil, maskAnyf(invalidConfigError, "HTTP handler must not be empty")
	}

	return newPrometheus, nil
}

type instrumentation struct {
	InstrumentationConfig

	Counters   map[string]spec.Counter
	Gauges     map[string]spec.Gauge
	Histograms map[string]spec.Histogram

	Mutex sync.Mutex
}

func (i *instrumentation) ExecFunc(key string, action func() error) error {
	h, err := i.GetHistogram(i.NewKey(key, "durations", "histogram", "milliseconds"))
	if err != nil {
		return maskAny(err)
	}
	c, err := i.GetCounter(i.NewKey(key, "errors", "counter", "total"))
	if err != nil {
		return maskAny(err)
	}

	start := time.Now()

	err = action()
	if err != nil {
		c.IncrBy(1)
		return maskAny(err)
	}

	stop := time.Now()
	sample := float64(stop.Sub(start).Nanoseconds() / 1000000)
	h.Observe(sample)

	return nil
}

func (i *instrumentation) GetCounter(key string) (spec.Counter, error) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	if c, ok := i.Counters[key]; ok {
		return c, nil
	}

	newConfig := DefaultCounterConfig()
	newConfig.Name = key
	newConfig.Help = helpFor("Counter", key)
	newCounter, err := NewCounter(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	err = prometheus.Register(newCounter.(*counter).ClientCounter)
	if err != nil {
		return nil, maskAny(err)
	}
	i.Counters[key] = newCounter

	return newCounter, nil
}

func (i *instrumentation) GetGauge(key string) (spec.Gauge, error) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	if g, ok := i.Gauges[key]; ok {
		return g, nil
	}

	newConfig := DefaultGaugeConfig()
	newConfig.Name = key
	newConfig.Help = helpFor("Gauge", key)
	newGauge, err := NewGauge(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	err = prometheus.Register(newGauge.(*gauge).ClientGauge)
	if err != nil {
		return nil, maskAny(err)
	}
	i.Gauges[key] = newGauge

	return newGauge, nil
}

func (i *instrumentation) GetHistogram(key string) (spec.Histogram, error) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()

	if h, ok := i.Histograms[key]; ok {
		return h, nil
	}

	newConfig := DefaultHistogramConfig()
	newConfig.Name = key
	newConfig.Help = helpFor("Histogram", key)
	newHistogram, err := NewHistogram(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	err = prometheus.Register(newHistogram.(*histogram).ClientHistogram)
	if err != nil {
		return nil, maskAny(err)
	}
	i.Histograms[key] = newHistogram

	return newHistogram, nil
}

func (i *instrumentation) GetHTTPEndpoint() string {
	return i.HTTPEndpoint
}

func (i *instrumentation) GetHTTPHandler() http.Handler {
	return i.HTTPHandler
}

func (i *instrumentation) GetPrefixes() []string {
	return i.Prefixes
}

func (i *instrumentation) NewKey(s ...string) string {
	return key.NewPromKey(append(i.Prefixes, s...)...)
}

func (i *instrumentation) WrapFunc(key string, action func() error) func() error {
	wrappedFunc := func() error {
		err := i.ExecFunc(key, action)
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	return wrappedFunc
}
