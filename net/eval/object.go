package evalnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (en *evalNet) GetID() spec.ObjectID {
	en.Mutex.Lock()
	defer en.Mutex.Unlock()
	return en.ID
}

func (en *evalNet) GetType() spec.ObjectType {
	en.Mutex.Lock()
	defer en.Mutex.Unlock()
	return en.Type
}
