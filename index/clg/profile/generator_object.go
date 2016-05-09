package profile

import (
	"github.com/xh3b4sd/anna/spec"
)

func (g *generator) GetID() spec.ObjectID {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	return g.ID
}

func (g *generator) GetType() spec.ObjectType {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	return g.Type
}
