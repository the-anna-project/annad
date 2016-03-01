package innet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (in *inNet) GetID() spec.ObjectID {
	in.Mutex.Lock()
	defer in.Mutex.Unlock()
	return in.ID
}

func (in *inNet) GetType() spec.ObjectType {
	in.Mutex.Lock()
	defer in.Mutex.Unlock()
	return in.Type
}
