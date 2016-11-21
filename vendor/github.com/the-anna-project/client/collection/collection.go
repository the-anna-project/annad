package collection

import (
	"sync"

	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new endpoint collection.
func New() servicespec.EndpointCollection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	metricService servicespec.EndpointService
	textService   servicespec.EndpointService

	// Settings.

	shutdownOnce sync.Once
}

func (c *collection) Boot() {
	//go c.Metric().Boot()
	go c.Text().Boot()
}

func (c *collection) Metric() servicespec.EndpointService {
	return c.metricService
}

func (c *collection) SetMetricService(metricService servicespec.EndpointService) {
	c.metricService = metricService
}

func (c *collection) SetTextService(textService servicespec.EndpointService) {
	c.textService = textService
}

func (c *collection) Shutdown() {
	c.shutdownOnce.Do(func() {
		var wg sync.WaitGroup

		//wg.Add(1)
		//go func() {
		//	c.Metric().Shutdown()
		//	wg.Done()
		//}()

		wg.Add(1)
		go func() {
			c.Text().Shutdown()
			wg.Done()
		}()

		wg.Wait()
	})
}

func (c *collection) Text() servicespec.EndpointService {
	return c.textService
}
