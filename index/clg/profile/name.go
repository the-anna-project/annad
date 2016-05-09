package profile

import (
	"github.com/xh3b4sd/anna/spec"
)

func (g *generator) CreateName(clgName string) (string, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call CreateName")

	return clgName, nil
}
