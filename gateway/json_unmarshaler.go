package gateway

import (
	"github.com/xh3b4sd/anna/spec"
)

func (g *gateway) UnmarshalJSON(bytes []byte) error {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 15}, "call UnmarshalJSON")
	return nil
}
