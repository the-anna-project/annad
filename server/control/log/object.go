package log

import (
	"github.com/xh3b4sd/anna/spec"
)

func (c *control) GetID() string {
	return c.ID
}

func (c *control) GetType() spec.ObjectType {
	return c.Type
}
