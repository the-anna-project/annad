package space

import (
	"github.com/xh3b4sd/anna/object/config/space/connection"
	"github.com/xh3b4sd/anna/object/config/space/dimension"
	"github.com/xh3b4sd/anna/object/config/space/peer"
)

// New creates a new space object. It provides configuration for the connection
// space.
func New() *Object {
	return &object{}
}

type Object struct {
	// Settings.

	connection *connection.Object
	dimension  *dimension.Object
	peer       *peer.Object
}

func (o *Object) Configure() error {
	// Settings.

	o.connection = connection.New()
	o.dimension = dimension.New()
	o.peer = peer.New()

	err := o.connection.Configure()
	if err != nil {
		return maskAny(err)
	}
	err = o.dimension.Configure()
	if err != nil {
		return maskAny(err)
	}
	err = o.peer.Configure()
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (o *Object) Connection() *connection.Object {
	return o.connection
}

func (o *Object) Dimension() *dimension.Object {
	return o.dimension
}

func (o *Object) Peer() *peer.Object {
	return o.peer
}

func (o *Object) Validate() error {
	return nil
}
