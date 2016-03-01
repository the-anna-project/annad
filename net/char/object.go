package charnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (cn *charNet) GetID() spec.ObjectID {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()
	return cn.ID
}

func (cn *charNet) GetType() spec.ObjectType {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()
	return cn.Type
}
