package impulse

import (
	"github.com/xh3b4sd/anna/spec"
)

func (i *impulse) GetID() spec.ObjectID {
	return i.ID
}

func (i *impulse) GetType() spec.ObjectType {
	return i.Type
}
