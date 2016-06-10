package strategy

import (
	"github.com/xh3b4sd/anna/spec"
)

func (d *dynamic) GetID() spec.ObjectID {
	return d.ID
}

func (d *dynamic) GetType() spec.ObjectType {
	return d.Type
}
