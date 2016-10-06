package network

import (
	"github.com/xh3b4sd/anna/spec"
)

func (n *network) Gateway() spec.GatewayCollection {
	return n.GatewayCollection
}
