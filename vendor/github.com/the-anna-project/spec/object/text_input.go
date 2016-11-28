package object

// TextInput represents a streamed request being send to the neural network.
// This is basically good for requesting calculations from the neural network
// by providing text input and an optional expectation object.
type TextInput interface {
	// Echo returns the echo flag of the current text request.
	Echo() bool
	// Expectation returns the expectation of the current text request.
	Expectation() Expectation
	// Input returns the input of the current text request.
	Input() string
	// SessionID returns the session ID of the current text request.
	SessionID() string
	SetEcho(echo bool)
	SetExpectation(expectation Expectation)
	SetInput(input string)
	SetSessionID(sessionID string)
}
