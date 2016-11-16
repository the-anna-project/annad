package endpoint

import (
	"sync"

	servicespec "github.com/the-anna-project/spec/service"
)

// NewCollection creates a new endpoint collection.
func NewCollection() servicespec.EndpointCollection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	metric servicespec.EndpointService
	text   servicespec.EndpointService

	// Settings.

	shutdownOnce sync.Once
}

func (c *collection) Boot() {
	go c.Metric().Boot()
	go c.Text().Boot()
}

func (c *collection) Metric() servicespec.EndpointService {
	return c.metric
}

func (c *collection) Text() servicespec.EndpointService {
	return c.text
}

func (c *collection) SetMetric(metric servicespec.EndpointService) {
	c.metric = metric
}

func (c *collection) SetText(text servicespec.EndpointService) {
	c.text = text
}

func (c *collection) Shutdown() {
	c.shutdownOnce.Do(func() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			c.Metric().Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			c.Text().Shutdown()
			wg.Done()
		}()

		wg.Wait()
	})
}
