package execnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (rn *execNet) GetID() spec.ObjectID {
	rn.Mutex.Lock()
	defer rn.Mutex.Unlock()
	return rn.ID
}

func (rn *execNet) GetType() spec.ObjectType {
	rn.Mutex.Lock()
	defer rn.Mutex.Unlock()
	return rn.Type
}
