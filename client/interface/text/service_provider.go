package text

import (
	"github.com/xh3b4sd/anna/spec"
)

func (c *client) Service() spec.ServiceCollection {
	return c.ServiceCollection
}
