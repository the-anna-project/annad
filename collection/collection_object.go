package clg

import (
	"github.com/xh3b4sd/anna/spec"
)

func (c *Collection) GetID() spec.ObjectID {
	return c.ID
}

func (c *Collection) GetType() spec.ObjectType {
	return c.Type
}
