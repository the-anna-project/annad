package memory

import (
	"net/http"

	"github.com/xh3b4sd/anna/spec"
)

// InstrumentationConfig represents the configuration used to create a new
// memory instrumentation object.
type InstrumentationConfig struct {
	// Settings.
}

// DefaultInstrumentationConfig provides a default configuration to create a
// new memory instrumentation object by best effort.
func DefaultInstrumentationConfig() InstrumentationConfig {
	newConfig := InstrumentationConfig{}

	return newConfig
}

// NewInstrumentation creates a new configured memory instrumentation object.
func NewInstrumentation(config InstrumentationConfig) (spec.Instrumentation, error) {
	newInstrumentation := &instrumentation{
		InstrumentationConfig: config,
	}

	return newInstrumentation, nil
}

type instrumentation struct {
	InstrumentationConfig
}

func (p *instrumentation) ExecFunc(key string, action func() error) error {
	err := action()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (p *instrumentation) GetCounter(key string) (spec.Counter, error) {
	newConfig := DefaultCounterConfig()
	newCounter, err := NewCounter(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newCounter, nil
}

func (p *instrumentation) GetGauge(key string) (spec.Gauge, error) {
	newConfig := DefaultGaugeConfig()
	newGauge, err := NewGauge(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newGauge, nil
}

func (p *instrumentation) GetHistogram(key string) (spec.Histogram, error) {
	newConfig := DefaultHistogramConfig()
	newHistogram, err := NewHistogram(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newHistogram, nil
}

func (p *instrumentation) GetHTTPEndpoint() string {
	return ""
}

func (p *instrumentation) GetHTTPHandler() http.Handler {
	return nil
}

func (p *instrumentation) GetPrefixes() []string {
	return nil
}

func (p *instrumentation) NewKey(s ...string) string {
	return ""
}

func (p *instrumentation) WrapFunc(key string, action func() error) func() error {
	wrappedFunc := func() error {
		err := p.ExecFunc(key, action)
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	return wrappedFunc
}
