package ideanet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (in *ideaNet) GetID() spec.ObjectID {
	in.Mutex.Lock()
	defer in.Mutex.Unlock()
	return in.ID
}

func (in *ideaNet) GetType() spec.ObjectType {
	in.Mutex.Lock()
	defer in.Mutex.Unlock()
	return in.Type
}
