package core

import (
	"github.com/xh3b4sd/anna/spec"
)

func (c *core) GetObjectID() spec.ObjectID {
	c.Log.V(11).Debugf("call Core.GetObjectID")

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	return c.State.GetObjectID()
}

func (c *core) GetObjectType() spec.ObjectType {
	c.Log.V(11).Debugf("call Core.GetObjectType")

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	return c.State.GetObjectType()
}

func (c *core) GetState() spec.State {
	c.Log.V(11).Debugf("call Core.GetState")

	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	return c.State
}

func (c *core) SetState(state spec.State) {
	c.Log.V(11).Debugf("call Core.SetState")

	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.State = state
}
