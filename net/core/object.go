package corenet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (cn *coreNet) GetID() spec.ObjectID {
	return cn.ID
}

func (cn *coreNet) GetType() spec.ObjectType {
	return cn.Type
}
