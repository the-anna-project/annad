package factoryserver

import (
	"github.com/xh3b4sd/anna/spec"
)

func (fs *factoryServer) GetID() spec.ObjectID {
	fs.Mutex.Lock()
	defer fs.Mutex.Unlock()
	return fs.ID
}

func (fs *factoryServer) GetType() spec.ObjectType {
	fs.Mutex.Lock()
	defer fs.Mutex.Unlock()
	return fs.Type
}
