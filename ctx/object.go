package ctx

import (
	"github.com/xh3b4sd/anna/spec"
)

func (c *ctx) GetID() spec.ObjectID {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	return c.ID
}

func (c *ctx) GetType() spec.ObjectType {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	return c.Type
}
