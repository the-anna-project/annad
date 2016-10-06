package gateway

import (
	"github.com/xh3b4sd/anna/gateway/text-output"
	"github.com/xh3b4sd/anna/spec"
)

// CollectionConfig represents the configuration used to create a new gateway
// collection object.
type CollectionConfig struct {
	// Dependencies.
	TextOutputGateway spec.TextOutputGateway
}

// DefaultCollectionConfig provides a default configuration to create a new
// gateway collection object by best effort.
func DefaultCollectionConfig() CollectionConfig {
	newConfig := CollectionConfig{
		// Dependencies.
		TextOutputGateway: textoutput.MustNewGateway(),
	}

	return newConfig
}

// NewCollection creates a new configured gateway collection object.
func NewCollection(config CollectionConfig) (spec.GatewayCollection, error) {
	newCollection := &collection{
		CollectionConfig: config,
	}

	if newCollection.TextOutputGateway == nil {
		return nil, maskAnyf(invalidConfigError, "text output gateway must not be empty")
	}

	return newCollection, nil
}

// MustNewCollection creates either a new default configured gateway collection
// object, or panics.
func MustNewCollection() spec.GatewayCollection {
	newCollection, err := NewCollection(DefaultCollectionConfig())
	if err != nil {
		panic(err)
	}

	return newCollection
}

type collection struct {
	CollectionConfig
}

func (c *collection) TextOutput() spec.TextOutputGateway {
	return c.TextOutputGateway
}
