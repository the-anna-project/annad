// Package expectation implements spec.Expectation to provide verification
// between input and out.
package expectation

import (
	"github.com/xh3b4sd/anna/object/spec"
)

// Config represents the configuration used to create a new expectation response
// object.
type Config struct {
	Output string
}

// DefaultConfig provides a default configuration to create a new expectation
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		Output: "",
	}

	return newConfig
}

// New creates a new configured expectation object.
func New(config Config) (spec.Expectation, error) {
	newExpectation := &expectation{
		Config: config,
	}

	return newExpectation, nil
}

type expectation struct {
	Config
}

// IsEmpty checks whether the current expectation is empty.
func (e *expectation) GetOutput() string {
	return e.Output
}
