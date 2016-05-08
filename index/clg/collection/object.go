package collection

import (
	"github.com/xh3b4sd/anna/spec"
)

func (c *collection) GetID() spec.ObjectID {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	return c.ID
}

func (c *collection) GetType() spec.ObjectType {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	return c.Type
}
