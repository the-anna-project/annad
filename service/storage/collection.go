package storage

import (
	"sync"

	"github.com/xh3b4sd/anna/service/spec"
)

// NewCollection creates a new storage collection.
func NewCollection() spec.StorageCollection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	connection spec.Storage
	feature    spec.Storage
	general    spec.Storage

	// Settings.

	shutdownOnce sync.Once
}

func (c *collection) Boot() {
	go c.Connection().Boot()
	go c.Feature().Boot()
	go c.General().Boot()
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

func (c *collection) SetConnection(conn spec.Storage) {
	c.connection = conn
}

func (c *collection) SetFeature(f spec.Storage) {
	c.feature = f
}

func (c *collection) SetGeneral(g spec.Storage) {
	c.general = g
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
