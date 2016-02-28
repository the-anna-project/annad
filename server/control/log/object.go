package logcontrol

import (
	"github.com/xh3b4sd/anna/spec"
)

func (lc *logControl) GetID() spec.ObjectID {
	lc.Mutex.Lock()
	defer lc.Mutex.Unlock()
	return lc.ID
}

func (lc *logControl) GetType() spec.ObjectType {
	lc.Mutex.Lock()
	defer lc.Mutex.Unlock()
	return lc.Type
}
