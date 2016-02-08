package server

import (
	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/spec"
)

// GetObjectID returns the hardcoded object ID 12cf5de8b0bd06e8.
func (s server) GetObjectID() spec.ObjectID {
	return spec.ObjectID("12cf5de8b0bd06e8")
}

func (s server) GetObjectType() spec.ObjectType {
	return common.ObjectType.Server
}

// GetState is not implemented for this object.
func (s server) GetState() spec.State {
	return nil
}

// SetState is not implemented for this object.
func (s server) SetState(state spec.State) {
}
