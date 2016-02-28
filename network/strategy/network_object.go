package strategynetwork

import (
	"github.com/xh3b4sd/anna/spec"
)

func (n *network) GetID() spec.ObjectID {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	return n.ID
}

func (n *network) GetType() spec.ObjectType {
	n.Mutex.Lock()
	defer n.Mutex.Unlock()
	return n.Type
}
