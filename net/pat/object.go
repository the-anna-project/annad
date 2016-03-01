package patnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (pn *patNet) GetID() spec.ObjectID {
	pn.Mutex.Lock()
	defer pn.Mutex.Unlock()
	return pn.ID
}

func (pn *patNet) GetType() spec.ObjectType {
	pn.Mutex.Lock()
	defer pn.Mutex.Unlock()
	return pn.Type
}
