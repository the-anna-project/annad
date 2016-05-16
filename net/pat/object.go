package patnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (pn *patNet) GetID() spec.ObjectID {
	return pn.ID
}

func (pn *patNet) GetType() spec.ObjectType {
	return pn.Type
}
