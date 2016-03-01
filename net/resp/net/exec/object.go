package execnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (en *execNet) GetID() spec.ObjectID {
	en.Mutex.Lock()
	defer en.Mutex.Unlock()
	return en.ID
}

func (en *execNet) GetType() spec.ObjectType {
	en.Mutex.Lock()
	defer en.Mutex.Unlock()
	return en.Type
}
