package config

import (
	"github.com/xh3b4sd/anna/object/config/endpoint"
	"github.com/xh3b4sd/anna/object/config/space"
	"github.com/xh3b4sd/anna/object/config/storage"
)

// New creates a new config collection. It provides configuration for the whole
// neural network.
func New() *Collection {
	return &Collection{}
}

type Collection struct {
	// Settings.

	endpointCollection *endpoint.Collection
	spaceCollection    *space.Collection
	storageCollection  *storage.Collection
}

func (c *Collection) Endpoint() *endpoint.Collection {
	return c.endpointCollection
}

func (c *Collection) SetEndpointCollection(endpointCollection *endpoint.Collection) {
	c.endpointCollection = endpointCollection
}

func (c *Collection) SetSpaceCollection(spaceCollection *space.Collection) {
	c.spaceCollection = spaceCollection
}

func (c *Collection) SetStorageCollection(storageCollection *storage.Collection) {
	c.storageCollection = storageCollection
}

func (c *Collection) Space() *space.Collection {
	return c.spaceCollection
}

func (c *Collection) Storage() *storage.Collection {
	return c.storageCollection
}
