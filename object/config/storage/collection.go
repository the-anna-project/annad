package storage

import (
	"github.com/xh3b4sd/anna/object/config/storage/connection"
	"github.com/xh3b4sd/anna/object/config/storage/feature"
	"github.com/xh3b4sd/anna/object/config/storage/general"
)

// NewCollection creates a new storage object. It provides configuration for the
// storage service.
func NewCollection() *Collection {
	return &Collection{}
}

// Collection represents the storage collection.
type Collection struct {
	// Settings.

	connection *connection.Object
	feature    *feature.Object
	general    *general.Object
}

// Connection returns the connection storage config of the storage collection.
func (c *Collection) Connection() *connection.Object {
	return c.connection
}

// Feature returns the feature storage config of the storage collection.
func (c *Collection) Feature() *feature.Object {
	return c.feature
}

// General returns the general storage config of the storage collection.
func (c *Collection) General() *general.Object {
	return c.general
}

// SetConnection sets the connection storage config for the storage collection.
func (c *Collection) SetConnection(connection *connection.Object) {
	c.connection = connection
}

// SetFeature sets the feature storage config for the storage collection.
func (c *Collection) SetFeature(feature *feature.Object) {
	c.feature = feature
}

// SetGeneral sets the general storage config for the storage collection.
func (c *Collection) SetGeneral(general *general.Object) {
	c.general = general
}
