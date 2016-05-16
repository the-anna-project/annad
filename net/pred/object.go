package prednet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (pn *predNet) GetID() spec.ObjectID {
	return pn.ID
}

func (pn *predNet) GetType() spec.ObjectType {
	return pn.Type
}
