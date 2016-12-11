// Package collection implements services to manage connections inside network
// layers.
package collection

import servicespec "github.com/the-anna-project/spec/service"

// New creates a new layer collection.
func New() servicespec.LayerCollection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	behaviourService   servicespec.LayerService
	informationService servicespec.LayerService
	positionService    servicespec.LayerService
}

func (c *collection) Boot() {
	go c.Behaviour().Boot()
	go c.Information().Boot()
	go c.Position().Boot()
}

func (c *collection) Behaviour() servicespec.LayerService {
	return c.behaviourService
}

func (c *collection) Information() servicespec.LayerService {
	return c.informationService
}

func (c *collection) Position() servicespec.LayerService {
	return c.positionService
}

func (c *collection) SetBehaviourService(behaviourService servicespec.LayerService) {
	c.behaviourService = behaviourService
}

func (c *collection) SetInformationService(informationService servicespec.LayerService) {
	c.informationService = informationService
}

func (c *collection) SetPositionService(positionService servicespec.LayerService) {
	c.positionService = positionService
}
