package spec

// TextEndpoint provides a way to feed neural networks with text input. It opens
// a stream endpoint for text requests which are forwarded to the neural
// network. Text response are streamed back to the client as soon as they are
// available. StreamText blocks until the given context is canceled.
type TextEndpoint interface {
	// Boot initializes and starts the text endpoint service like booting a
	// machine. The call to Boot runs a gRPC server and blocks forever, so you
	// might want to call it in a separate goroutine.
	Boot()

	Configure() error

	Service() Collection

	SetGRPCAddress(gRPCAddr string)

	SetServiceCollection(sc Collection)

	// Shutdown ends all processes of the text endpoint service like shutting down
	// a machine. The call to Shutdown blocks until the text endpoint service is
	// completely shut down, so you might want to call it in a separate goroutine.
	Shutdown()

	Validate() error
}
