package space

import (
	"github.com/the-anna-project/annad/object/config/space/connection"
	"github.com/the-anna-project/annad/object/config/space/dimension"
	"github.com/the-anna-project/annad/object/config/space/peer"
)

// NewCollection creates a new space object. It provides configuration for the
// connection space.
func NewCollection() *Collection {
	return &Collection{}
}

// Collection represents the space collection.
type Collection struct {
	// Settings.

	connection *connection.Object
	dimension  *dimension.Object
	peer       *peer.Object
}

// Connection returns the connection space config of the space collection.
func (c *Collection) Connection() *connection.Object {
	return c.connection
}

// Dimension returns the space dimension config of the space collection.
func (c *Collection) Dimension() *dimension.Object {
	return c.dimension
}

// Peer returns the space peer config of the space collection.
func (c *Collection) Peer() *peer.Object {
	return c.peer
}

// SetConnection sets the connection space config for the space collection.
func (c *Collection) SetConnection(connection *connection.Object) {
	c.connection = connection
}

// SetDimension sets the space dimension config for the space collection.
func (c *Collection) SetDimension(dimension *dimension.Object) {
	c.dimension = dimension
}

// SetPeer sets the space peer config for the space collection.
func (c *Collection) SetPeer(peer *peer.Object) {
	c.peer = peer
}
