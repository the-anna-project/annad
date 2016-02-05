package core

import (
	"github.com/xh3b4sd/anna/spec"
)

func (c *core) GetObjectID() spec.ObjectID {
	return c.State.GetObjectID()
}

func (c *core) GetObjectType() spec.ObjectType {
	return c.State.GetObjectType()
}

func (c *core) GetState() spec.State {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	return c.State
}

func (c *core) SetState(state spec.State) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.State = state
}
