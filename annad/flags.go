package main

// Flags represents the flags of the command line object.
type Flags struct {
	GRPCAddr string
	HTTPAddr string

	ControlLogLevels    string
	ControlLogObejcts   string
	ControlLogVerbosity int

	Storage                 string
	RedisFeatureStorageAddr string
	RedisGeneralStorageAddr string
	RedisStoragePrefix      string
}
