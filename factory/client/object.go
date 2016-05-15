package factoryclient

import (
	"github.com/xh3b4sd/anna/spec"
)

func (fc *factoryClient) GetID() spec.ObjectID {
	return fc.ID
}

func (fc *factoryClient) GetType() spec.ObjectType {
	return fc.Type
}
