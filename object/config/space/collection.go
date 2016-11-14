package space

import (
	"github.com/xh3b4sd/anna/object/config/space/connection"
	"github.com/xh3b4sd/anna/object/config/space/dimension"
	"github.com/xh3b4sd/anna/object/config/space/peer"
)

// New creates a new space object. It provides configuration for the connection
// space.
func New() *Collection {
	return &Collection{}
}

type Collection struct {
	// Settings.

	connection *connection.Object
	dimension  *dimension.Object
	peer       *peer.Object
}

func (c *Collection) Connection() *connection.Object {
	return c.connection
}

func (c *Collection) Dimension() *dimension.Object {
	return c.dimension
}

func (c *Collection) Peer() *peer.Object {
	return c.peer
}
