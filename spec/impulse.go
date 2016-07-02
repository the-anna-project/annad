package spec

// Impulse is basically a container to carry around information. It walks
// through neural networks and is modified on its way to finally creating some
// output based on some input.
//
// An impulse carries information about former impulses and their corresponding
// inputs around. This enables the network to collect related information to
// make a reasonable point about the current contextual task.
//
type Impulse interface {
	// GetExpectation returns the currently configured expectation object.
	GetExpectation() Expectation

	// GetInput returns the currently configured input.
	GetInput() string

	// GetOutput returns the impulse's output.
	GetOutput() string

	// GetSessionID returns the ID of the session related to the current impulse.
	GetSessionID() string

	Object

	// SetOutput sets the impulse's output.
	SetOutput(output string)
}
