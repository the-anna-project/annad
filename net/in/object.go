package innet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (in *inNet) GetID() spec.ObjectID {
	return in.ID
}

func (in *inNet) GetType() spec.ObjectType {
	return in.Type
}
