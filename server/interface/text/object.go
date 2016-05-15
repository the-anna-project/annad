package textinterface

import (
	"github.com/xh3b4sd/anna/spec"
)

func (ti *textInterface) GetID() spec.ObjectID {
	return ti.ID
}

func (ti *textInterface) GetType() spec.ObjectType {
	return ti.Type
}
