package profile

import (
	"github.com/xh3b4sd/anna/spec"
)

func (p *profile) GetID() spec.ObjectID {
	return p.ID
}

func (p *profile) GetType() spec.ObjectType {
	return p.Type
}
