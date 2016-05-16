// Package prometheus implements spec.Instrumentation and provides
// instrumentation primitives to manage application metrics.
package prometheus

import (
	"net/http"
	"sync"
	"time"

	prometheusclient "github.com/prometheus/client_golang/prometheus"

	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
)

// Config represents the configuration used to create a new prometheus
// instrumentation object.
type Config struct {
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

// DefaultConfig provides a default configuration to create a new prometheus
// instrumentation object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// Settings.
		Prefixes:     []string{"anna"},
		HTTPEndpoint: "/metrics",
		HTTPHandler:  prometheusclient.Handler(),
	}

	return newConfig
}

// New creates a new configured prometheus instrumentation object.
func New(config Config) (spec.Instrumentation, error) {
	newPrometheus := &prometheus{
		Config: config,

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

type prometheus struct {
	Config

	Counters   map[string]spec.Counter
	Gauges     map[string]spec.Gauge
	Histograms map[string]spec.Histogram

	Mutex sync.Mutex
}

func (p *prometheus) ExecFunc(key string, action func() error) error {
	h, err := p.GetHistogram(p.NewKey(key, "durations", "histogram", "milliseconds"))
	if err != nil {
		return maskAny(err)
	}
	c, err := p.GetCounter(p.NewKey(key, "errors", "counter", "total"))
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

func (p *prometheus) GetCounter(key string) (spec.Counter, error) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	if c, ok := p.Counters[key]; ok {
		return c, nil
	}

	newConfig := DefaultCounterConfig()
	newConfig.Name = key
	// TODO configure help
	newConfig.Help = "help"
	newCounter, err := NewCounter(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}
	_, err = prometheusclient.RegisterOrGet(newCounter.(*counter).ClientCounter)
	if err != nil {
		return nil, maskAny(err)
	}
	p.Counters[key] = newCounter

	return newCounter, nil
}

func (p *prometheus) GetGauge(key string) (spec.Gauge, error) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	if g, ok := p.Gauges[key]; ok {
		return g, nil
	}

	newConfig := DefaultGaugeConfig()
	newConfig.Name = key
	// TODO configure help
	newConfig.Help = "help"
	newGauge, err := NewGauge(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}
	_, err = prometheusclient.RegisterOrGet(newGauge.(*gauge).ClientGauge)
	if err != nil {
		return nil, maskAny(err)
	}
	p.Gauges[key] = newGauge

	return newGauge, nil
}

func (p *prometheus) GetHistogram(key string) (spec.Histogram, error) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	if h, ok := p.Histograms[key]; ok {
		return h, nil
	}

	newConfig := DefaultHistogramConfig()
	newConfig.Name = key
	// TODO configure help
	newConfig.Help = "help"
	newHistogram, err := NewHistogram(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}
	_, err = prometheusclient.RegisterOrGet(newHistogram.(*histogram).ClientHistogram)
	if err != nil {
		return nil, maskAny(err)
	}
	p.Histograms[key] = newHistogram

	return newHistogram, nil
}

func (p *prometheus) GetHTTPEndpoint() string {
	return p.HTTPEndpoint
}

func (p *prometheus) GetHTTPHandler() http.Handler {
	return p.HTTPHandler
}

func (p *prometheus) GetPrefixes() []string {
	return p.Prefixes
}

func (p *prometheus) NewKey(s ...string) string {
	return key.NewPromKey(append(p.Prefixes, s...)...)
}

func (p *prometheus) WrapFunc(key string, action func() error) func() error {
	wrappedFunc := func() error {
		err := p.ExecFunc(key, action)
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	return wrappedFunc
}
