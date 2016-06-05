package stratnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (sn *stratNet) GetID() spec.ObjectID {
	return sn.ID
}

func (sn *stratNet) GetType() spec.ObjectType {
	return sn.Type
}
