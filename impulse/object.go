package impulse

import (
	"github.com/xh3b4sd/anna/spec"
)

func (i *impulse) GetID() spec.ObjectID {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	return i.ID
}

func (i *impulse) GetType() spec.ObjectType {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	return i.Type
}
