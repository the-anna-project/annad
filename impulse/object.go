package impulse

import (
	"github.com/xh3b4sd/anna/spec"
)

func (i *impulse) GetObjectID() spec.ObjectID {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	return i.State.GetObjectID()
}

func (i *impulse) GetObjectType() spec.ObjectType {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	return i.State.GetObjectType()
}

func (i *impulse) GetState() spec.State {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	return i.State
}

func (i *impulse) SetState(state spec.State) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	i.State = state
}
