package osfilesystem

import (
	"github.com/xh3b4sd/anna/spec"
)

func (osfs *osFileSystem) GetID() spec.ObjectID {
	osfs.Mutex.Lock()
	defer osfs.Mutex.Unlock()
	return osfs.ID
}

func (osfs *osFileSystem) GetType() spec.ObjectType {
	osfs.Mutex.Lock()
	defer osfs.Mutex.Unlock()
	return osfs.Type
}
