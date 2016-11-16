package spec

// TextInput represents a streamed request being send to the neural network.
// This is basically good for requesting calculations from the neural network
// by providing text input and an optional expectation object.
type TextInput interface {
	// GetEcho returns the echo flag of the current text request.
	GetEcho() bool

	// GetExpectation returns the expectation of the current text request.
	GetExpectation() Expectation

	// GetInput returns the input of the current text request.
	GetInput() string

	// GetSessionID returns the session ID of the current text request.
	GetSessionID() string
}
