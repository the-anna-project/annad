package network

import (
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

func (n *network) Service() servicespec.Collection {
	return n.ServiceCollection
}
