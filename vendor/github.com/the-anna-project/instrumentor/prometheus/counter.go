package prometheus

import (
	prometheusclient "github.com/prometheus/client_golang/prometheus"

	objectspec "github.com/the-anna-project/spec/object"
)

// CounterConfig represents the configuration used to create a new prometheus
// counter object.
type CounterConfig struct {
	// Settings.

	// Help represents some sort of informative description of the registered
	// metric.
	Help string

	// Name represents the metric's key as it is supposed to be registered. In the
	// scope of prometheus this is expected to be an underscored string.
	Name string
}

// DefaultCounterConfig provides a default configuration to create a new
// prometheus counter object by best effort.
func DefaultCounterConfig() CounterConfig {
	newConfig := CounterConfig{
		// Settings.
		Help: "",
		Name: "",
	}

	return newConfig
}

// NewCounter creates a new configured prometheus counter object.
func NewCounter(config CounterConfig) (objectspec.InstrumentorCounter, error) {
	newCounter := &counter{
		CounterConfig: config,
	}

	if newCounter.Help == "" {
		return nil, maskAnyf(invalidConfigError, "help must not be empty")
	}
	if newCounter.Name == "" {
		return nil, maskAnyf(invalidConfigError, "name must not be empty")
	}

	newCounter.ClientCounter = prometheusclient.NewCounter(prometheusclient.CounterOpts{
		Help: newCounter.Help,
		Name: newCounter.Name,
	})

	return newCounter, nil
}

type counter struct {
	CounterConfig

	ClientCounter prometheusclient.Counter
}

func (c *counter) IncrBy(delta float64) {
	c.ClientCounter.Add(delta)
}
