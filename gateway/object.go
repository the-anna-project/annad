package gateway

import (
	"github.com/xh3b4sd/anna/spec"
)

func (g *gateway) GetID() spec.ObjectID {
	return g.ID
}

func (g *gateway) GetType() spec.ObjectType {
	return g.Type
}
