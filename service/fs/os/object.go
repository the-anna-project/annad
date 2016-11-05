package os

import (
	"github.com/xh3b4sd/anna/spec"
)

func (osfs *osFileSystem) GetID() string {
	return osfs.ID
}

func (osfs *osFileSystem) GetType() spec.ObjectType {
	return osfs.Type
}
