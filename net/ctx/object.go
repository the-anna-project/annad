package ctxnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (cn *ctxNet) GetID() spec.ObjectID {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()
	return cn.ID
}

func (cn *ctxNet) GetType() spec.ObjectType {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()
	return cn.Type
}
