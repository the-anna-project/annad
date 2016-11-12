package prometheus

import (
	prometheusclient "github.com/prometheus/client_golang/prometheus"

	"github.com/xh3b4sd/anna/service/spec"
)

// GaugeConfig represents the configuration used to create a new prometheus
// gauge object.
type GaugeConfig struct {
	// Settings.

	// Help represents some sort of informative description of the registered
	// metric.
	Help string

	// Name represents the metric's key as it is supposed to be registered. In the
	// scope of prometheus this is expected to be an underscored string.
	Name string
}

// DefaultGaugeConfig provides a default configuration to create a new
// prometheus gauge object by best effort.
func DefaultGaugeConfig() GaugeConfig {
	newConfig := GaugeConfig{
		// Settings.
		Help: "",
		Name: "",
	}

	return newConfig
}

// NewGauge creates a new configured prometheus gauge object.
func NewGauge(config GaugeConfig) (spec.Gauge, error) {
	newGauge := &gauge{
		GaugeConfig: config,
	}

	if newGauge.Help == "" {
		return nil, maskAnyf(invalidConfigError, "help must not be empty")
	}
	if newGauge.Name == "" {
		return nil, maskAnyf(invalidConfigError, "name must not be empty")
	}

	newGauge.ClientGauge = prometheusclient.NewGauge(prometheusclient.GaugeOpts{
		Help: newGauge.Help,
		Name: newGauge.Name,
	})

	return newGauge, nil
}

type gauge struct {
	GaugeConfig

	ClientGauge prometheusclient.Gauge
}

func (g *gauge) DecrBy(delta float64) {
	g.ClientGauge.Sub(delta)
}

func (g *gauge) IncrBy(delta float64) {
	g.ClientGauge.Add(delta)
}
