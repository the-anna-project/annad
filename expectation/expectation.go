// Package expectation implements spec.Expectation to provide verification
// between input and out.
package expectation

import (
	"github.com/xh3b4sd/anna/spec"
)

// ExpectationConfig represents the configuration used to create a new
// expectation response object.
type ExpectationConfig struct {
	Output string
}

// DefaultExpectationConfig provides a default configuration to create a new
// expectation object by best effort.
func DefaultExpectationConfig() ExpectationConfig {
	newConfig := ExpectationConfig{
		Output: "",
	}

	return newConfig
}

// NewExpectation creates a new configured expectation object.
func NewExpectation(config ExpectationConfig) (spec.Expectation, error) {
	newExpectation := &expectation{
		ExpectationConfig: config,
	}

	return newExpectation, nil
}

// Expectation represents the payload of an spec.Expectation.
type expectation struct {
	ExpectationConfig
}

// IsEmpty checks whether the current expectation is empty.
func (e *expectation) GetOutput() string {
	return er.Output
}
