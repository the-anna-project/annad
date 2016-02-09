package logcontrol

import (
	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/spec"
)

// GetObjectID returns the hardcoded object ID 126c88f053e7eff2.
func (lc logControl) GetObjectID() spec.ObjectID {
	return spec.ObjectID("126c88f053e7eff2")
}

func (lc logControl) GetObjectType() spec.ObjectType {
	return common.ObjectType.LogControl
}

// GetState is not implemented for this object.
func (lc logControl) GetState() spec.State {
	return nil
}

// SetState is not implemented for this object.
func (lc logControl) SetState(state spec.State) {
}
