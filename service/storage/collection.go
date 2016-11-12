package storage

import (
	"sync"

	"github.com/xh3b4sd/anna/service/spec"
)

// NewCollection creates a new storage collection service.
func NewCollection() spec.StorageCollection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	connection        spec.Storage
	feature           spec.Storage
	general           spec.Storage
	serviceCollection spec.Collection

	// Settings.

	metadata     map[string]string
	shutdownOnce sync.Once
}

func (c *collection) Configure() error {
	// Settings.

	id, err := c.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}
	c.metadata = map[string]string{
		"id":   id,
		"kind": "storage",
		"name": "collection",
		"type": "service",
	}

	c.shutdownOnce = sync.Once{}

	return nil
}

func (c *collection) Connection() spec.Storage {
	return c.connection
}

func (c *collection) Feature() spec.Storage {
	return c.feature
}

func (c *collection) General() spec.Storage {
	return c.general
}

func (c *collection) Metadata() map[string]string {
	return c.metadata
}

func (c *collection) Service() spec.Collection {
	return c.serviceCollection
}

func (c *collection) SetConnection(conn spec.Storage) {
	c.connection = conn
}

func (c *collection) SetFeature(f spec.Storage) {
	c.feature = f
}

func (c *collection) SetGeneral(g spec.Storage) {
	c.general = g
}

func (c *collection) SetServiceCollection(sc spec.Collection) {
	c.serviceCollection = sc
}

func (c *collection) Shutdown() {
	c.shutdownOnce.Do(func() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			c.Connection().Shutdown()
			wg.Done()
		}()

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

func (c *collection) Validate() error {
	// Dependencies.

	if c.connection == nil {
		return maskAnyf(invalidConfigError, "connection storage must not be empty")
	}
	if c.feature == nil {
		return maskAnyf(invalidConfigError, "feature storage must not be empty")
	}
	if c.general == nil {
		return maskAnyf(invalidConfigError, "general storage must not be empty")
	}
	if c.serviceCollection == nil {
		return maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	return nil
}
