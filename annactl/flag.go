package main

// InterfaceTextReadPlainFlags represents the flags of the interface text read
// plain command.
type InterfaceTextReadPlainFlags struct {
	// Expectation represents the expectation object in JSON format.
	Expectation string
}

// Flags represents the flags of the command line object.
type Flags struct {
	Addr                string
	ControlLogLevels    string
	ControlLogVerbosity int

	InterfaceTextReadPlain InterfaceTextReadPlainFlags
}
