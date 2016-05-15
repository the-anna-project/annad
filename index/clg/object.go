package clg

import (
	"github.com/xh3b4sd/anna/spec"
)

func (i *index) GetID() spec.ObjectID {
	return i.ID
}

func (i *index) GetType() spec.ObjectType {
	return i.Type
}
