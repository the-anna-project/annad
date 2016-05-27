package memory

import (
	"github.com/xh3b4sd/anna/spec"
)

// HistogramConfig represents the configuration used to create a new memory
// histogram object.
type HistogramConfig struct {
}

// DefaultHistogramConfig provides a default configuration to create a new
// memory histogram object by best effort.
func DefaultHistogramConfig() HistogramConfig {
	newConfig := HistogramConfig{}

	return newConfig
}

// NewHistogram creates a new configured memory histogram object.
func NewHistogram(config HistogramConfig) (spec.Histogram, error) {
	newHistogram := &histogram{
		HistogramConfig: config,
	}

	return newHistogram, nil
}

type histogram struct {
	HistogramConfig
}

func (h *histogram) Observe(sample float64) {
}
