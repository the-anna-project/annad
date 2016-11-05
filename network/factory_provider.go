package network

import (
	"github.com/xh3b4sd/anna/spec"
)

func (n *network) Service() spec.ServiceCollection {
	return n.ServiceCollection
}
