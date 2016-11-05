package network

import (
	"github.com/xh3b4sd/anna/spec"
)

func (n *network) GetID() string {
	return n.ID
}

func (n *network) GetType() spec.ObjectType {
	return n.Type
}
