package spec

import (
	"encoding/json"

	"golang.org/x/net/context"
)

// TextRequest represents a streamed request being send to the neural network.
// This is basically good for requesting calculations from the neural network
// by providing text input and an optional expectation object.
type TextRequest interface {
	// GetInput returns the input of the current text request.
	GetInput() string

	json.Unmarshaler

	// IsEmpty checks whether the current text request is empty. An empty text
	// request can be considered invalid.
	IsEmpty() bool
}

// TextResponse represents a streamed response being send to the client. This
// is basically good for responding calculated output of the neural network.
type TextResponse interface {
	// GetOutput returns the output of the current text response.
	GetOutput() string

	json.Unmarshaler
}

// TextInterface provides a way to feed neural networks with text input.
type TextInterface interface {
	// StreamText forwards the text request provided by in to the neural network
	// and forwards the text response to the client. StreamText blocks until the
	// given context is canceled.
	StreamText(ctx context.Context, in chan TextRequest, out chan TextResponse) error
}
