package factory

import (
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/factory/permutation"
	"github.com/xh3b4sd/anna/factory/random"
	"github.com/xh3b4sd/anna/spec"
)

// CollectionConfig represents the configuration used to create a new factory
// collection object.
type CollectionConfig struct {
	// Dependencies.
	IDFactory          spec.IDFactory
	PermutationFactory spec.PermutationFactory
	RandomFactory      spec.RandomFactory
}

// DefaultCollectionConfig provides a default configuration to create a new
// factory collection object by best effort.
func DefaultCollectionConfig() CollectionConfig {
	newConfig := CollectionConfig{
		// Dependencies.
		IDFactory:          id.MustNewFactory(),
		PermutationFactory: permutation.MustNewFactory(),
		RandomFactory:      random.MustNewFactory(),
	}

	return newConfig
}

// NewCollection creates a new configured factory collection object.
func NewCollection(config CollectionConfig) (spec.FactoryCollection, error) {
	newCollection := &collection{
		CollectionConfig: config,
	}

	if newCollection.IDFactory == nil {
		return nil, maskAnyf(invalidConfigError, "ID factory must not be empty")
	}
	if newCollection.PermutationFactory == nil {
		return nil, maskAnyf(invalidConfigError, "permutation factory must not be empty")
	}
	if newCollection.RandomFactory == nil {
		return nil, maskAnyf(invalidConfigError, "random factory must not be empty")
	}

	return newCollection, nil
}

// MustNewCollection creates either a new default configured factory collection
// object, or panics.
func MustNewCollection() spec.FactoryCollection {
	newCollection, err := NewCollection(DefaultCollectionConfig())
	if err != nil {
		panic(err)
	}

	return newCollection
}

type collection struct {
	CollectionConfig
}

func (c *collection) ID() spec.IDFactory {
	return c.IDFactory
}

func (c *collection) Permutation() spec.PermutationFactory {
	return c.PermutationFactory
}

func (c *collection) Random() spec.RandomFactory {
	return c.RandomFactory
}
