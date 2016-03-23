package patnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (d *distribution) GetID() spec.ObjectID {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	return d.ID
}

func (d *distribution) GetType() spec.ObjectType {
	d.Mutex.Lock()
	defer d.Mutex.Unlock()
	return d.Type
}
