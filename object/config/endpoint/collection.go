package endpoint

import (
	"github.com/xh3b4sd/anna/object/config/endpoint/metric"
	"github.com/xh3b4sd/anna/object/config/endpoint/text"
)

// NewCollection creates a new endpoint object. It provides configuration for
// the network endpoints.
func NewCollection() *Collection {
	return &Collection{}
}

// Collection represents the endpoint collection.
type Collection struct {
	// Settings.

	metric *metric.Object
	text   *text.Object
}

// Metric returns the metric config of the endpoint collection.
func (c *Collection) Metric() *metric.Object {
	return c.metric
}

// Text returns the text config of the endpoint collection.
func (c *Collection) Text() *text.Object {
	return c.text
}

// SetMetric sets the metric config for the endpoint collection.
func (c *Collection) SetMetric(metric *metric.Object) {
	c.metric = metric
}

// SetText sets the text config for the endpoint collection.
func (c *Collection) SetText(text *text.Object) {
	c.text = text
}
