package endpoint

import (
	"github.com/xh3b4sd/anna/object/config/endpoint/metric"
	"github.com/xh3b4sd/anna/object/config/endpoint/text"
)

// New creates a new endpoint object. It provides configuration for the network
// endpoints.
func New() *Collection {
	return &Collection{}
}

type Collection struct {
	// Settings.

	metric *metric.Object
	text   *text.Object
}

func (c *Collection) Metric() *metric.Object {
	return c.metric
}

func (c *Collection) Text() *text.Object {
	return c.text
}
