package core

import (
	"github.com/xh3b4sd/anna/spec"
)

func (c *core) GetID() spec.ObjectID {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	return c.ID
}

func (c *core) GetType() spec.ObjectType {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	return c.Type
}
