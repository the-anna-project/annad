package log

import (
	"github.com/xh3b4sd/anna/spec"
)

func (c *control) GetID() spec.ObjectID {
	return c.ID
}

func (c *control) GetType() spec.ObjectType {
	return c.Type
}
