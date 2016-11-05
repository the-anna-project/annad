package spec

import (
	"github.com/xh3b4sd/anna/service/spec"
	"golang.org/x/net/context"
)

// ExpectationRequest represents an expectation given with a text request to
// define the expected outcome of a requested calculation. That way the
// calculated result of the neural network can be checked against something the
// client expects to be returned. By providing an expectation request the neural
// network can be trained, because it calculates and iterates over its own
// activities until the calculated output equals the provided expectation.
type ExpectationRequest interface {
	// GetExpectation returns the expectation of the current expectation request.
	GetExpectation() Expectation
}

// TextRequest represents a streamed request being send to the neural network.
// This is basically good for requesting calculations from the neural network
// by providing text input and an optional expectation object.
type TextRequest interface {
	// GetEcho returns the echo flag of the current text request.
	GetEcho() bool

	// GetExpectation returns the expectation of the current text request.
	GetExpectation() Expectation

	// GetInput returns the input of the current text request.
	GetInput() string

	// GetSessionID returns the session ID of the current text request.
	GetSessionID() string
}

// TextInterfaceClient provides a way to feed neural networks with text input.
type TextInterfaceClient interface {
	// StreamText forwards the text request provided by in to the neural network
	// and forwards the text response to the client. StreamText blocks until the
	// given context is canceled.
	StreamText(ctx context.Context, in chan TextRequest, out chan spec.TextResponse) error
}
