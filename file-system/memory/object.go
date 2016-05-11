package memoryfilesystem

import (
	"github.com/xh3b4sd/anna/spec"
)

func (mfs *memoryFileSystem) GetID() spec.ObjectID {
	mfs.Mutex.Lock()
	defer mfs.Mutex.Unlock()
	return mfs.ID
}

func (mfs *memoryFileSystem) GetType() spec.ObjectType {
	mfs.Mutex.Lock()
	defer mfs.Mutex.Unlock()
	return mfs.Type
}
