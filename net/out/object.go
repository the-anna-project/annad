package outnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (on *outNet) GetID() spec.ObjectID {
	on.Mutex.Lock()
	defer on.Mutex.Unlock()
	return on.ID
}

func (on *outNet) GetType() spec.ObjectType {
	on.Mutex.Lock()
	defer on.Mutex.Unlock()
	return on.Type
}
