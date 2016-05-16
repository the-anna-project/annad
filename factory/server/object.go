package factoryserver

import (
	"github.com/xh3b4sd/anna/spec"
)

func (fs *factoryServer) GetID() spec.ObjectID {
	return fs.ID
}

func (fs *factoryServer) GetType() spec.ObjectType {
	return fs.Type
}
