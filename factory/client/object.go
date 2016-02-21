package factoryclient

import (
	"github.com/xh3b4sd/anna/spec"
)

func (fc *factoryClient) GetID() spec.ObjectID {
	fc.Mutex.Lock()
	defer fc.Mutex.Unlock()
	return fc.ID
}

func (fc *factoryClient) GetType() spec.ObjectType {
	fc.Mutex.Lock()
	defer fc.Mutex.Unlock()
	return fc.Type
}
