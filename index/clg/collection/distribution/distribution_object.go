package distribution

import (
	"github.com/xh3b4sd/anna/spec"
)

func (d *distribution) GetID() string {
	return d.ID
}

func (d *distribution) GetType() spec.ObjectType {
	return d.Type
}
