package stratnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (sn *stratNet) GetID() spec.ObjectID {
	sn.Mutex.Lock()
	defer sn.Mutex.Unlock()
	return sn.ID
}

func (sn *stratNet) GetType() spec.ObjectType {
	sn.Mutex.Lock()
	defer sn.Mutex.Unlock()
	return sn.Type
}
