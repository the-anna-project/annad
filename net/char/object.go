package charnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (cn *charNet) GetID() spec.ObjectID {
	return cn.ID
}

func (cn *charNet) GetType() spec.ObjectType {
	return cn.Type
}
