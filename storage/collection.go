package storage

import (
	"sync"

	"github.com/xh3b4sd/anna/storage/memory"
	"github.com/xh3b4sd/anna/storage/spec"
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

// NewCollection creates a new configured storage collection object.
func NewCollection(config CollectionConfig) (spec.Collection, error) {
	newCollection := &collection{
		CollectionConfig: config,

		ShutdownOnce: sync.Once{},
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
func MustNewCollection() spec.Collection {
	newCollection, err := NewCollection(DefaultCollectionConfig())
	if err != nil {
		panic(err)
	}

	return newCollection
}

type collection struct {
	CollectionConfig

	ShutdownOnce sync.Once
}

func (c *collection) Feature() spec.Storage {
	return c.FeatureStorage
}

func (c *collection) General() spec.Storage {
	return c.GeneralStorage
}

func (c *collection) Shutdown() {
	c.ShutdownOnce.Do(func() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			c.Feature().Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			c.General().Shutdown()
			wg.Done()
		}()

		wg.Wait()
	})
}
