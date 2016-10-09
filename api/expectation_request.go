package api

import (
	"github.com/xh3b4sd/anna/spec"
)

// ExpectationRequestConfig represents the configuration used to create a new
// expectation response object.
type ExpectationRequestConfig struct {
	Expectation spec.Expectation
}

// DefaultExpectationRequestConfig provides a default configuration to create a
// new expectation request object by best effort.
func DefaultExpectationRequestConfig() ExpectationRequestConfig {
	newConfig := ExpectationRequestConfig{
		Expectation: nil,
	}

	return newConfig
}

// NewExpectationRequest creates a new configured expectation request object.
func NewExpectationRequest(config ExpectationRequestConfig) (spec.ExpectationRequest, error) {
	newExpectationRequest := &expectationRequest{
		ExpectationRequestConfig: config,
	}

	return newExpectationRequest, nil
}

// ExpectationRequest represents the request payload of an spec.Expectation.
type expectationRequest struct {
	ExpectationRequestConfig
}

// IsEmpty checks whether the current expectation is empty.
func (er *expectationRequest) GetExpectation() spec.Expectation {
	return er.Expectation
}
