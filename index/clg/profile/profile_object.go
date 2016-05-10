package profile

import (
	"github.com/xh3b4sd/anna/spec"
)

func (p *profile) GetID() spec.ObjectID {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	return p.ID
}

func (p *profile) GetType() spec.ObjectType {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	return p.Type
}
