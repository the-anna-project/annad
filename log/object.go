package log

import (
	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/spec"
)

// GetObjectID returns the hardcoded object ID 126c88f053e7eff2.
func (l *log) GetObjectID() spec.ObjectID {
	return spec.ObjectID("126c88f053e7eff2")
}

func (l *log) GetObjectType() spec.ObjectType {
	return common.ObjectType.Log
}

// GetState is not implemented for this object.
func (l *log) GetState() spec.State {
	return nil
}

// SetState is not implemented for this object.
func (l *log) SetState(state spec.State) {
}
