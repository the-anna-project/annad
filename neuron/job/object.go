package jobneuron

import (
	"github.com/xh3b4sd/anna/spec"
)

func (n *neuron) GetObjectID() spec.ObjectID {
	return n.State.GetObjectID()
}

func (n *neuron) GetObjectType() spec.ObjectType {
	return n.State.GetObjectType()
}

func (n *neuron) GetState() spec.State {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	return n.State
}

func (n *neuron) SetState(state spec.State) {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	n.State = state
}
