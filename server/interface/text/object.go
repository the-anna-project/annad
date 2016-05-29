package text

import (
	"github.com/xh3b4sd/anna/spec"
)

func (i *tinterface) GetID() spec.ObjectID {
	return i.ID
}

func (i *tinterface) GetType() spec.ObjectType {
	return i.Type
}
