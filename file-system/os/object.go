package osfilesystem

import (
	"github.com/xh3b4sd/anna/spec"
)

func (osfs *osFileSystem) GetID() spec.ObjectID {
	return osfs.ID
}

func (osfs *osFileSystem) GetType() spec.ObjectType {
	return osfs.Type
}
