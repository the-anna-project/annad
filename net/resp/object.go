package respnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (rn *respNet) GetID() spec.ObjectID {
	return rn.ID
}

func (rn *respNet) GetType() spec.ObjectType {
	return rn.Type
}
