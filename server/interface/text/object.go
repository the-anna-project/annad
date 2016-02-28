package textinterface

import (
	"github.com/xh3b4sd/anna/spec"
)

func (ti *textInterface) GetID() spec.ObjectID {
	ti.Mutex.Lock()
	defer ti.Mutex.Unlock()
	return ti.ID
}

func (ti *textInterface) GetType() spec.ObjectType {
	ti.Mutex.Lock()
	defer ti.Mutex.Unlock()
	return ti.Type
}
