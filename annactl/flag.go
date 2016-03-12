package main

type InterfaceTextReadPlainFlags struct {
	Expected string
}

type Flags struct {
	Addr                string
	ControlLogLevels    string
	ControlLogVerbosity int

	InterfaceTextReadPlain InterfaceTextReadPlainFlags
}
