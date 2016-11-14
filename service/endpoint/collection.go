package endpoint

import (
	"sync"

	"github.com/xh3b4sd/anna/service/spec"
)

// NewCollection creates a new endpoint collection.
func NewCollection() spec.EndpointCollection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	metric spec.Endpoint
	text   spec.Endpoint

	// Settings.

	shutdownOnce sync.Once
}

func (c *collection) Boot() {
	go c.Metric().Boot()
	go c.Text().Boot()
}

func (c *collection) Metric() spec.Endpoint {
	return c.metric
}

func (c *collection) Text() spec.Endpoint {
	return c.text
}

func (c *collection) SetMetric(metric spec.Endpoint) {
	c.metric = metric
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
