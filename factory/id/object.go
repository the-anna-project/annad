package id

import (
	"github.com/xh3b4sd/anna/spec"
)

func (f *factory) GetID() spec.ObjectID {
	return f.ID
}

func (f *factory) GetType() spec.ObjectType {
	return f.Type
}
