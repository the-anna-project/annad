package api

import (
	"github.com/xh3b4sd/anna/service/spec"
)

// TextRequestConfig represents the configuration used to create a new text
// response object.
type TextRequestConfig struct {
	// Settings.

	// Echo being set to true causes the provided input simply to be echoed back.
	// The provided input goes through the whole stack and is streamed back and
	// forth, but bypasses neural network. This is useful to test the
	// client/server integration of the gRPC stream implementation.
	Echo bool

	// ExpectationRequest represents the expectation object. This is used to
	// match against the calculated output. In case there is an expectation
	// given, the neural network tries to calculate an output until it generated
	// one that matches the given expectation.
	ExpectationRequest spec.ExpectationRequest

	// Input represents the input being fed into the neural network. There must
	// be a none empty input given when requesting calculations from the neural
	// network.
	Input string

	// SessionID represents the session the current text request is associated
	// with. This is provided to differentiate streams between different users.
	SessionID string
}

// DefaultTextRequestConfig provides a default configuration to create a new
// text request object by best effort.
func DefaultTextRequestConfig() TextRequestConfig {
	newConfig := TextRequestConfig{
		Echo:               false,
		ExpectationRequest: nil,
		Input:              "",
		SessionID:          "",
	}

	return newConfig
}

// NewTextRequest creates a new configured text request object.
func NewTextRequest(config TextRequestConfig) (spec.TextRequest, error) {
	newTextRequest := &textRequest{
		TextRequestConfig: config,
	}

	if newTextRequest.Input == "" {
		return nil, maskAnyf(invalidConfigError, "input must not be empty")
	}
	if newTextRequest.SessionID == "" {
		return nil, maskAnyf(invalidConfigError, "session ID must not be empty")
	}

	return newTextRequest, nil
}

// MustNewTextRequest creates either a new default configured text request
// object, or panics.
func MustNewTextRequest() spec.TextRequest {
	newTextRequest, err := NewTextRequest(DefaultTextRequestConfig())
	if err != nil {
		panic(err)
	}

	return newTextRequest
}

type textRequest struct {
	TextRequestConfig
}

func (tr *textRequest) GetEcho() bool {
	return tr.Echo
}

func (tr *textRequest) GetExpectation() spec.Expectation {
	return tr.ExpectationRequest.GetExpectation()
}

func (tr *textRequest) GetInput() string {
	return tr.Input
}

func (tr *textRequest) GetSessionID() string {
	return tr.SessionID
}
