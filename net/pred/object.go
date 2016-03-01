package prednet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (pn *predNet) GetID() spec.ObjectID {
	pn.Mutex.Lock()
	defer pn.Mutex.Unlock()
	return pn.ID
}

func (pn *predNet) GetType() spec.ObjectType {
	pn.Mutex.Lock()
	defer pn.Mutex.Unlock()
	return pn.Type
}
