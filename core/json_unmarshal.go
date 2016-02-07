package core

import (
	"github.com/xh3b4sd/anna/spec"
)

func (c *core) UnmarshalJSON(bytes []byte) error {
	c.Log.WithTags(spec.Tags{L: "D", O: c, T: nil, V: 13}, "call UnmarshalJSON")

	err := c.GetState().SetStateFromObjectBytes(bytes)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
