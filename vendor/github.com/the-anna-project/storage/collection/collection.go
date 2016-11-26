// Package collection implements services to persist data. The storage
// collection bundles storage instances to pass them around.
package collection

import (
	"sync"

	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new storage collection.
func New() servicespec.StorageCollection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	connectionService servicespec.StorageService
	featureService    servicespec.StorageService
	generalService    servicespec.StorageService

	// Settings.

	shutdownOnce sync.Once
}

func (c *collection) Boot() {
	go c.Connection().Boot()
	go c.Feature().Boot()
	go c.General().Boot()
}

func (c *collection) Connection() servicespec.StorageService {
	return c.connectionService
}

func (c *collection) Feature() servicespec.StorageService {
	return c.featureService
}

func (c *collection) General() servicespec.StorageService {
	return c.generalService
}

func (c *collection) SetConnection(connectionService servicespec.StorageService) {
	c.connectionService = connectionService
}

func (c *collection) SetFeature(featureService servicespec.StorageService) {
	c.featureService = featureService
}

func (c *collection) SetGeneral(generalService servicespec.StorageService) {
	c.generalService = generalService
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
