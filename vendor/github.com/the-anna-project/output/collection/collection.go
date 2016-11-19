package collection

import (
	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new output collection.
func New() servicespec.OutputCollection {
	return &collection{}
}

type collection struct {
	// Dependencies.

	textService servicespec.OutputService
}

func (c *collection) Boot() {
	go c.Text().Boot()
}

func (c *collection) SetTextService(textService servicespec.OutputService) {
	c.textService = textService
}

func (c *collection) Text() servicespec.OutputService {
	return c.textService
}
