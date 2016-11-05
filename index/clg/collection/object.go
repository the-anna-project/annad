package collection

import (
	"github.com/xh3b4sd/anna/spec"
)

func (c *collection) GetID() string {
	return c.ID
}

func (c *collection) GetType() spec.ObjectType {
	return c.Type
}
