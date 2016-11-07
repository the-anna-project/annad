package network

import (
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

func (n *network) Storage() storagespec.Collection {
	return n.StorageCollection
}
