package execnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (rn *execNet) GetID() spec.ObjectID {
	return rn.ID
}

func (rn *execNet) GetType() spec.ObjectType {
	return rn.Type
}
