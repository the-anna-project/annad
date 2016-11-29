package storage

import (
	"github.com/the-anna-project/annad/object/config/storage/connection"
	"github.com/the-anna-project/annad/object/config/storage/feature"
	"github.com/the-anna-project/annad/object/config/storage/general"
	"github.com/the-anna-project/annad/object/config/storage/peer"
)

// NewCollection creates a new storage object. It provides configuration for the
// storage service.
func NewCollection() *Collection {
	return &Collection{}
}

// Collection represents the storage collection.
type Collection struct {
	// Settings.

	connectionObject *connection.Object
	featureObject    *feature.Object
	generalObject    *general.Object
	peerObject       *peer.Object
}

// Connection returns the connection storage config of the storage collection.
func (c *Collection) Connection() *connection.Object {
	return c.connectionObject
}

// Feature returns the feature storage config of the storage collection.
func (c *Collection) Feature() *feature.Object {
	return c.featureObject
}

// General returns the general storage config of the storage collection.
func (c *Collection) General() *general.Object {
	return c.generalObject
}

// Peer returns the peer storage config of the storage collection.
func (c *Collection) Peer() *peer.Object {
	return c.peerObject
}

// SetConnection sets the connection storage config for the storage collection.
func (c *Collection) SetConnection(connectionObject *connection.Object) {
	c.connectionObject = connectionObject
}

// SetFeature sets the feature storage config for the storage collection.
func (c *Collection) SetFeature(featureObject *feature.Object) {
	c.featureObject = featureObject
}

// SetGeneral sets the general storage config for the storage collection.
func (c *Collection) SetGeneral(generalObject *general.Object) {
	c.generalObject = generalObject
}

// SetPeer sets the peer storage config for the storage collection.
func (c *Collection) SetPeer(peerObject *peer.Object) {
	c.peerObject = peerObject
}
