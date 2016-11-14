package spec

import (
	"golang.org/x/net/context"

	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// TextInterfaceClient provides a way to feed neural networks with text input.
type TextInterfaceClient interface {
	Boot()

	Service() servicespec.Collection

	SetGRPCAddress(gRPCAddr string)

	SetServiceCollection(sc servicespec.Collection)

	// StreamText forwards the text request provided by in to the neural network
	// and forwards the text response to the client. StreamText blocks until the
	// given context is canceled.
	StreamText(ctx context.Context) error
}
