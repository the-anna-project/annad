package corenet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (cn *coreNet) GetID() spec.ObjectID {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()
	return cn.ID
}

func (cn *coreNet) GetType() spec.ObjectType {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()
	return cn.Type
}
