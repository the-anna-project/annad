package respnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (rn *respNet) GetID() spec.ObjectID {
	rn.Mutex.Lock()
	defer rn.Mutex.Unlock()
	return rn.ID
}

func (rn *respNet) GetType() spec.ObjectType {
	rn.Mutex.Lock()
	defer rn.Mutex.Unlock()
	return rn.Type
}
