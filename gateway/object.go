package gateway

import (
	"github.com/xh3b4sd/anna/spec"
)

func (g *gateway) GetID() spec.ObjectID {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	return g.ID
}

func (g *gateway) GetType() spec.ObjectType {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	return g.Type
}
