package memoryfilesystem

import (
	"github.com/xh3b4sd/anna/spec"
)

func (mfs *memoryFileSystem) GetID() spec.ObjectID {
	return mfs.ID
}

func (mfs *memoryFileSystem) GetType() spec.ObjectType {
	return mfs.Type
}
