package text

import (
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

func (c *client) Service() servicespec.Collection {
	return c.ServiceCollection
}
