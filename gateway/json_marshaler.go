package gateway

import (
	"github.com/xh3b4sd/anna/spec"
)

func (g *gateway) MarshalJSON() ([]byte, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 15}, "call MarshalJSON")
	return nil, nil
}
