package ideanet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (in *ideaNet) GetID() spec.ObjectID {
	return in.ID
}

func (in *ideaNet) GetType() spec.ObjectType {
	return in.Type
}
