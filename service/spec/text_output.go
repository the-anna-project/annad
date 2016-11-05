package spec

// TextOutput provides a communication channel to send information sequences
// back to the client.
type TextOutput interface {
	// GetChannel returns a channel which is used to send text responses back to
	// the client.
	GetChannel() chan TextResponse
}

// TextResponse represents a streamed response being send to the client. This
// is basically good for responding calculated output of the neural network.
type TextResponse interface {
	// GetOutput returns the output of the current text response.
	GetOutput() string
}
