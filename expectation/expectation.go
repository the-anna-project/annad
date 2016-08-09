// Package expectation implements spec.Expectation to provide verification
// between input and out.
package expectation

import (
	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

// Config represents the configuration used to create a new expectation object.
type Config struct {
	ExpectationRequest api.ExpectationRequest
}

// DefaultConfig provides a default configuration to create a new expectation
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		ExpectationRequest: api.ExpectationRequest{},
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

func (e *expectation) IsEmpty() bool {
	return e.ExpectationRequest.IsEmpty()
}

// TODO
func (e *expectation) Match() (bool, error) {
	return false, nil
}
