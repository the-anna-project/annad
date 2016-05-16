package outnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (on *outNet) GetID() spec.ObjectID {
	return on.ID
}

func (on *outNet) GetType() spec.ObjectType {
	return on.Type
}
