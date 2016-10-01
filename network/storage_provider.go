package network

import (
	"github.com/xh3b4sd/anna/spec"
)

func (n *network) Storage() spec.StorageCollection {
	return n.StorageCollection
}
