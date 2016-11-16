package spec

import (
	"golang.org/x/net/context"

	servicespec "github.com/the-anna-project/spec/service"
)

// TextInterfaceClient provides a way to feed neural networks with text input.
type TextInterfaceClient interface {
	Boot()

	Service() servicespec.ServiceCollection

	SetGRPCAddress(gRPCAddr string)

	SetServiceCollection(serviceCollection servicespec.ServiceCollection)

	// StreamText forwards the text request provided by in to the neural network
	// and forwards the text response to the client. StreamText blocks until the
	// given context is canceled.
	StreamText(ctx context.Context) error
}
