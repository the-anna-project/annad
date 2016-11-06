package spec

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
	GetExpectation() Expectation

	// GetInput returns the input of the current text request.
	GetInput() string

	// GetSessionID returns the session ID of the current text request.
	GetSessionID() string
}
