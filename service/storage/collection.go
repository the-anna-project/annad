package storage

import (
	"sync"

	servicespec "github.com/the-anna-project/spec/service"
)

// NewCollection creates a new storage collection.
func NewCollection() servicespec.StorageCollection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	connection servicespec.StorageService
	feature    servicespec.StorageService
	general    servicespec.StorageService

	// Settings.

	shutdownOnce sync.Once
}

func (c *collection) Boot() {
	go c.Connection().Boot()
	go c.Feature().Boot()
	go c.General().Boot()
}

func (c *collection) Connection() servicespec.StorageService {
	return c.connection
}

func (c *collection) Feature() servicespec.StorageService {
	return c.feature
}

func (c *collection) General() servicespec.StorageService {
	return c.general
}

func (c *collection) SetConnection(conn servicespec.StorageService) {
	c.connection = conn
}

func (c *collection) SetFeature(f servicespec.StorageService) {
	c.feature = f
}

func (c *collection) SetGeneral(g servicespec.StorageService) {
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
