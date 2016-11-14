// Package memory mocks spec.Instrumentor and does effectively nothing. It
// is only used for default configurations that require a satisfied
// instrumentation implementation. This should then be overwritten with a valid
// implementation if required.
package memory

import (
	"net/http"

	"github.com/xh3b4sd/anna/spec"
)

// New creates a new memory instrumentation service.
func New() spec.Instrumentor {
	return &service{}
}

type service struct {
}

func (s *service) Boot() {
}

func (s *service) ExecFunc(key string, action func() error) error {
	err := action()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) GetCounter(key string) (spec.Counter, error) {
	newConfig := DefaultCounterConfig()
	newCounter, err := NewCounter(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newCounter, nil
}

func (s *service) GetGauge(key string) (spec.Gauge, error) {
	newConfig := DefaultGaugeConfig()
	newGauge, err := NewGauge(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newGauge, nil
}

func (s *service) GetHistogram(key string) (spec.Histogram, error) {
	newConfig := DefaultHistogramConfig()
	newHistogram, err := NewHistogram(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newHistogram, nil
}

func (s *service) GetHTTPEndpoint() string {
	return ""
}

func (s *service) GetHTTPHandler() http.Handler {
	return nil
}

func (s *service) GetPrefixes() []string {
	return nil
}

func (s *service) Metadata() map[string]string {
	return nil
}

func (s *service) NewKey(s ...string) string {
	return ""
}

func (s *service) Service() servicespec.Collection {
	return nil
}

func (s *service) SetServiceCollection(sc servicespec.Collection) {
}

func (s *service) WrapFunc(key string, action func() error) func() error {
	wrappedFunc := func() error {
		err := s.ExecFunc(key, action)
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	return wrappedFunc
}
