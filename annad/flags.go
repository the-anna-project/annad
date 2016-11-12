package main

// Flags represents the flags of the command line object.
type Flags struct {
	GRPCAddr string
	HTTPAddr string

	// TODO better nesting/structuring of storage flags
	StorageType string

	RedisConnectionStorageAddr   string
	RedisConnectionStoragePrefix string

	RedisFeatureStorageAddr   string
	RedisFeatureStoragePrefix string

	RedisGeneralStorageAddr   string
	RedisGeneralStoragePrefix string
}
