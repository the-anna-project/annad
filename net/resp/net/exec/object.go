package execnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (en *execNet) GetID() spec.ObjectID {
	return en.ID
}

func (en *execNet) GetType() spec.ObjectType {
	return en.Type
}
