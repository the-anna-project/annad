package textinterface

import (
	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/spec"
)

// GetObjectID returns the hardcoded object ID 794e6df6087d04c1.
func (ti textInterface) GetObjectID() spec.ObjectID {
	return spec.ObjectID("794e6df6087d04c1")
}

func (ti textInterface) GetObjectType() spec.ObjectType {
	return common.ObjectType.TextInterface
}

// GetState is not implemented for this object.
func (ti textInterface) GetState() spec.State {
	return nil
}

// SetState is not implemented for this object.
func (ti textInterface) SetState(state spec.State) {
}
