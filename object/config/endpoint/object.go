package endpoint

import (
	"github.com/xh3b4sd/anna/object/config/endpoint/metric"
	"github.com/xh3b4sd/anna/object/config/endpoint/text"
)

// New creates a new endpoint object. It provides configuration for the network
// endpoints.
func New() *Object {
	return &object{}
}

type Object struct {
	// Settings.

	metric *metric.Object
	text   *text.Object
}

func (o *Object) Configure() error {
	// Settings.

	o.metric = metric.New()
	o.text = text.New()

	err := o.metric.Configure()
	if err != nil {
		return maskAny(err)
	}
	err = o.text.Configure()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (o *Object) Metric() *metric.Object {
	return o.metric
}

func (o *Object) Text() *text.Object {
	return o.text
}

func (o *Object) Validate() error {
	return nil
}
