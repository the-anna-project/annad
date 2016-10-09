package network

import (
	"github.com/xh3b4sd/anna/spec"
)

func (n *network) Factory() spec.FactoryCollection {
	return n.FactoryCollection
}
