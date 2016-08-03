package prometheus

import (
	prometheusclient "github.com/prometheus/client_golang/prometheus"

	"github.com/xh3b4sd/anna/spec"
)

// HistogramConfig represents the configuration used to create a new prometheus
// histogram object.
type HistogramConfig struct {
	// Settings.

	// Buckets represents a list of time ranges in seconds. Observed samples are
	// put into their corresponding ranges.
	//
	// A bucket's unit MUST be second. The buckets list MUST be ordered
	// incrementally.
	//
	// The buckets need to be properly configured to match the use case of the
	// oberseved samples, otherwise the histogram becomes pretty useless. E.g.
	// mapping samples of 25 milliseconds into a 5 second bucket makes no sense.
	Buckets []float64

	// Help represents some sort of informative description of the registered
	// metric.
	Help string

	// Name represents the metric's key as it is supposed to be registered. In the
	// scope of prometheus this is expected to be an underscored string.
	Name string
}

// DefaultHistogramConfig provides a default configuration to create a new
// prometheus histogram object by best effort.
func DefaultHistogramConfig() HistogramConfig {
	newConfig := HistogramConfig{
		// Settings.
		Buckets: []float64{.001, .002, .003, .004, .005, .01, .02, .03, .04, .05, .1, .2, .3, .4, .5, 1, 2, 3, 4, 5, 10},
		Help:    "",
		Name:    "",
	}

	return newConfig
}

// NewHistogram creates a new configured prometheus histogram object.
func NewHistogram(config HistogramConfig) (spec.Histogram, error) {
	newHistogram := &histogram{
		HistogramConfig: config,
	}

	if len(newHistogram.Buckets) == 0 {
		return nil, maskAnyf(invalidConfigError, "buckets must not be empty")
	}
	if newHistogram.Help == "" {
		return nil, maskAnyf(invalidConfigError, "help must not be empty")
	}
	if newHistogram.Name == "" {
		return nil, maskAnyf(invalidConfigError, "name must not be empty")
	}

	newHistogram.ClientHistogram = prometheusclient.NewHistogram(prometheusclient.HistogramOpts{
		Buckets: newHistogram.Buckets,
		Help:    newHistogram.Help,
		Name:    newHistogram.Name,
	})

	return newHistogram, nil
}

type histogram struct {
	HistogramConfig

	ClientHistogram prometheusclient.Histogram
}

func (h *histogram) Observe(sample float64) {
	h.ClientHistogram.Observe(sample)
}
