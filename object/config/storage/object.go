package storage

import (
	"github.com/xh3b4sd/anna/object/config/storage/connection"
	"github.com/xh3b4sd/anna/object/config/storage/feature"
	"github.com/xh3b4sd/anna/object/config/storage/general"
)

// New creates a new storage object. It provides configuration for the storage
// service.
func New() *Object {
	return &object{}
}

type Object struct {
	// Settings.

	connection *connection.Object
	feature    *feature.Object
	general    *general.Object
}

func (o *Object) Configure() error {
	// Settings.

	o.connection = connection.New()
	o.feature = feature.New()
	o.general = general.New()

	err := o.connection.Configure()
	if err != nil {
		return maskAny(err)
	}
	err = o.feature.Configure()
	if err != nil {
		return maskAny(err)
	}
	err = o.general.Configure()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (o *Object) Connection() *connection.Object {
	return o.connection
}

func (o *Object) Feature() *feature.Object {
	return o.feature
}

func (o *Object) General() *general.Object {
	return o.general
}

func (o *Object) Validate() error {
	return nil
}
