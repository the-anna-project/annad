package storage

import (
	"github.com/xh3b4sd/anna/object/config/storage/connection"
	"github.com/xh3b4sd/anna/object/config/storage/feature"
	"github.com/xh3b4sd/anna/object/config/storage/general"
)

// New creates a new storage object. It provides configuration for the storage
// service.
func New() *Collection {
	return &Collection{}
}

type Collection struct {
	// Settings.

	connection *connection.Object
	feature    *feature.Object
	general    *general.Object
}

func (c *Collection) Connection() *connection.Object {
	return c.connection
}

func (c *Collection) Feature() *feature.Object {
	return c.feature
}

func (c *Collection) General() *general.Object {
	return c.general
}
