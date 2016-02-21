package core

import (
	"github.com/xh3b4sd/anna/spec"
)

func (c *core) MarshalJSON() ([]byte, error) {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 15}, "call MarshalJSON")
	return nil, nil
}
