package config

import (
	"github.com/xh3b4sd/anna/object/config/endpoint"
	"github.com/xh3b4sd/anna/object/config/space"
	"github.com/xh3b4sd/anna/object/config/storage"
)

// NewCollection creates a new config collection. It provides configuration for
// the whole neural network.
func NewCollection() *Collection {
	return &Collection{}
}

// Collection represents the config collection.
type Collection struct {
	// Settings.

	endpointCollection *endpoint.Collection
	spaceCollection    *space.Collection
	storageCollection  *storage.Collection
}

// Endpoint returns the endpoint collection of the config collection.
func (c *Collection) Endpoint() *endpoint.Collection {
	return c.endpointCollection
}

// SetEndpointCollection sets the endpoint collection for the config collection.
func (c *Collection) SetEndpointCollection(endpointCollection *endpoint.Collection) {
	c.endpointCollection = endpointCollection
}

// SetSpaceCollection sets the space collection for the config collection.
func (c *Collection) SetSpaceCollection(spaceCollection *space.Collection) {
	c.spaceCollection = spaceCollection
}

// SetStorageCollection sets the storage collection for the config collection.
func (c *Collection) SetStorageCollection(storageCollection *storage.Collection) {
	c.storageCollection = storageCollection
}

// Space returns the space collection of the config collection.
func (c *Collection) Space() *space.Collection {
	return c.spaceCollection
}

// Storage returns the storage collection of the config collection.
func (c *Collection) Storage() *storage.Collection {
	return c.storageCollection
}
