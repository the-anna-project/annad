package network

import (
	"github.com/xh3b4sd/anna/spec"
)

func (n *network) GetObjectID() spec.ObjectID {
	return n.State.GetObjectID()
}

func (n *network) GetObjectType() spec.ObjectType {
	return n.State.GetObjectType()
}

func (n *network) GetState() spec.State {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	return n.State
}

func (n *network) SetState(state spec.State) {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	n.State = state
}
