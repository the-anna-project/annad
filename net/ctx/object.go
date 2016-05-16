package ctxnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (cn *ctxNet) GetID() spec.ObjectID {
	return cn.ID
}

func (cn *ctxNet) GetType() spec.ObjectType {
	return cn.Type
}
