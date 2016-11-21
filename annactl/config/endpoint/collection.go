package endpoint

import "github.com/xh3b4sd/anna/annactl/config/endpoint/text"

// NewCollection creates a new endpoint object. It provides configuration for
// the network endpoints.
func NewCollection() *Collection {
	return &Collection{}
}

// Collection represents the endpoint collection.
type Collection struct {
	// Settings.

	textObject *text.Object
}

// Text returns the text config of the endpoint collection.
func (c *Collection) Text() *text.Object {
	return c.textObject
}

// SetText sets the text config for the endpoint collection.
func (c *Collection) SetText(textObject *text.Object) {
	c.textObject = textObject
}
