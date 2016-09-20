package storage

import (
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

// CollectionConfig represents the configuration used to create a new storage
// collection object.
type CollectionConfig struct {
	// Dependencies.
	FeatureStorage spec.Storage
	GeneralStorage spec.Storage
}

// DefaultCollectionConfig provides a default configuration to create a new
// storage collection object by best effort.
func DefaultCollectionConfig() CollectionConfig {
	newConfig := CollectionConfig{
		// Dependencies.
		FeatureStorage: memory.MustNew(),
		GeneralStorage: memory.MustNew(),
	}

	return newConfig
}

// NewCollection creates a new configured network object.
func NewCollection(config CollectionConfig) (spec.StorageCollection, error) {
	newCollection := &collection{
		CollectionConfig: config,
	}

	if newCollection.FeatureStorage == nil {
		return nil, maskAnyf(invalidConfigError, "feature storage must not be empty")
	}
	if newCollection.GeneralStorage == nil {
		return nil, maskAnyf(invalidConfigError, "general storage must not be empty")
	}

	return newCollection, nil
}

// MustNewCollection creates either a new default configured storage collection
// object, or panics.
func MustNewCollection() spec.StorageCollection {
	newCollection, err := NewCollection(DefaultCollectionConfig())
	if err != nil {
		panic(err)
	}

	return newCollection
}

type collection struct {
	CollectionConfig
}

func (c *collection) Feature() spec.Storage {
	return c.FeatureStorage
}

func (c *collection) General() spec.Storage {
	return c.GeneralStorage
}
