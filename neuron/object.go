package neuron

import (
	"github.com/xh3b4sd/anna/spec"
)

//
// character
//

func (cn *characterNeuron) GetObjectID() spec.ObjectID {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()
	return cn.State.GetObjectID()
}

func (cn *characterNeuron) GetObjectType() spec.ObjectType {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()
	return cn.State.GetObjectType()
}

func (cn *characterNeuron) GetState() spec.State {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()
	return cn.State
}

func (cn *characterNeuron) SetState(state spec.State) {
	cn.Mutex.Lock()
	defer cn.Mutex.Unlock()
	cn.State = state
}

//
// first
//

func (fn *firstNeuron) GetObjectID() spec.ObjectID {
	fn.Mutex.Lock()
	defer fn.Mutex.Unlock()
	return fn.State.GetObjectID()
}

func (fn *firstNeuron) GetObjectType() spec.ObjectType {
	fn.Mutex.Lock()
	defer fn.Mutex.Unlock()
	return fn.State.GetObjectType()
}

func (fn *firstNeuron) GetState() spec.State {
	fn.Mutex.Lock()
	defer fn.Mutex.Unlock()
	return fn.State
}

func (fn *firstNeuron) SetState(state spec.State) {
	fn.Mutex.Lock()
	defer fn.Mutex.Unlock()
	fn.State = state
}

//
// job
//

func (jn *jobNeuron) GetObjectID() spec.ObjectID {
	jn.Mutex.Lock()
	defer jn.Mutex.Unlock()
	return jn.State.GetObjectID()
}

func (jn *jobNeuron) GetObjectType() spec.ObjectType {
	jn.Mutex.Lock()
	defer jn.Mutex.Unlock()
	return jn.State.GetObjectType()
}

func (jn *jobNeuron) GetState() spec.State {
	jn.Mutex.Lock()
	defer jn.Mutex.Unlock()
	return jn.State
}

func (jn *jobNeuron) SetState(state spec.State) {
	jn.Mutex.Lock()
	defer jn.Mutex.Unlock()
	jn.State = state
}
