package main

// InterfaceTextReadPlainFlags represents the flags of the interface text read
// plain command.
type InterfaceTextReadPlainFlags struct {
	Expected string
}

// Flags represents the flags of the command line object.
type Flags struct {
	Addr                string
	ControlLogLevels    string
	ControlLogVerbosity int

	InterfaceTextReadPlain InterfaceTextReadPlainFlags
}
