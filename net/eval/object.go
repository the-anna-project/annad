package evalnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (en *evalNet) GetID() spec.ObjectID {
	return en.ID
}

func (en *evalNet) GetType() spec.ObjectType {
	return en.Type
}
