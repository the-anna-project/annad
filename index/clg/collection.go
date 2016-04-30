package clg

import (
	"github.com/xh3b4sd/anna/spec"
)

// CollectionConfig represents the configuration used to create a new CLG
// collection object.
type CollectionConfig struct{}

// DefaultCLGCollectionConfig provides a default configuration to create a new
// CLG collection object by best effort.
func DefaultCLGCollectionConfig() CollectionConfig {
	newConfig := CollectionConfig{}

	return newConfig
}

// NewCLGCollection creates a new configured CLG collection object.
func NewCLGCollection(config CollectionConfig) (spec.CLGCollection, error) {
	newCLGCollection := &clgCollection{
		CollectionConfig: config,
	}

	return newCLGCollection, nil
}

type clgCollection struct {
	CollectionConfig
}
