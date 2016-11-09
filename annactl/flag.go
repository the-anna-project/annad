package main

// InterfaceTextReadPlainFlags represents the flags of the interface text read
// plain command.
type InterfaceTextReadPlainFlags struct {
	// Echo being set to true causes the provided input simply to be echoed back.
	// The provided input goes through the whole stack and is streamed back and
	// forth, but bypasses neural network. This is useful to test the
	// client/server integration of the gRPC stream implementation.
	Echo bool

	// Expectation represents the expectation object in JSON format.
	Expectation string
}

// Flags represents the flags of the command line object.
type Flags struct {
	GRPCAddr string
	HTTPAddr string

	InterfaceTextReadPlain InterfaceTextReadPlainFlags
}
