package collection

import (
	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new input collection.
func New() servicespec.InputCollection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	textService servicespec.InputService
}

func (c *collection) Boot() {
	go c.Text().Boot()
}

func (c *collection) SetTextService(textService servicespec.InputService) {
	c.textService = textService
}

func (c *collection) Text() servicespec.InputService {
	return c.textService
}
