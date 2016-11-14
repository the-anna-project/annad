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

func (o *Collection) Endpoint() *endpoint.Collection {
	return o.endpointCollection
}

func (o *Collection) SetEndpointCollection(endpointCollection *endpoint.Collection) {
	o.endpointCollection = endpointCollection
}

func (o *Collection) SetSpaceCollection(spaceCollection *space.Collection) {
	o.spaceCollection = spaceCollection
}

func (o *Collection) SetStorageCollection(storageCollection *storage.Collection) {
	o.storageCollection = storageCollection
}

func (o *Collection) Space() *space.Collection {
	return o.spaceCollection
}

func (o *Collection) Storage() *storage.Collection {
	return o.storageCollection
}
