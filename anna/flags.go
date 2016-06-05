package main

// Flags represents the flags of the command line object.
type Flags struct {
	Addr string

	ControlLogLevels    string
	ControlLogObejcts   string
	ControlLogVerbosity int

	Storage     string
	StorageAddr string
}
