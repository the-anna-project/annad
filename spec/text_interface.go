package spec

import (
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
)

// TextInterface provides a way to feed neural networks with text input.
type TextInterface interface {
	// StreamText forwards the text request provided by in to the neural network
	// and forwards the text response to the client. StreamText blocks until the
	// given context is canceled.
	StreamText(ctx context.Context, in chan api.TextRequest, out chan api.TextResponse) error
}
