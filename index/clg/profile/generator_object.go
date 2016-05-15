package profile

import (
	"github.com/xh3b4sd/anna/spec"
)

func (g *generator) GetID() spec.ObjectID {
	return g.ID
}

func (g *generator) GetType() spec.ObjectType {
	return g.Type
}
