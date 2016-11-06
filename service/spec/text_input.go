package spec

import "github.com/xh3b4sd/anna/object/spec"

// TextInput provides a communication channel to send information sequences
// back to the client.
type TextInput interface {
	// GetChannel returns a channel which is used to send text responses back to
	// the client.
	GetChannel() chan TextRequest
}

// TextRequest represents a streamed request being send to the neural network.
// This is basically good for requesting calculations from the neural network
// by providing text input and an optional expectation object.
type TextRequest interface {
	// GetEcho returns the echo flag of the current text request.
	GetEcho() bool

	// GetExpectation returns the expectation of the current text request.
	GetExpectation() spec.Expectation

	// GetInput returns the input of the current text request.
	GetInput() string

	// GetSessionID returns the session ID of the current text request.
	GetSessionID() string
}
