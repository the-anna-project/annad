package spec

import (
	"github.com/xh3b4sd/anna/service/spec"
	"golang.org/x/net/context"
)

// TextInterfaceClient provides a way to feed neural networks with text input.
type TextInterfaceClient interface {
	// StreamText forwards the text request provided by in to the neural network
	// and forwards the text response to the client. StreamText blocks until the
	// given context is canceled.
	// TODO text interface should use service collection
	StreamText(ctx context.Context, in chan spec.TextRequest, out chan spec.TextResponse) error
}
