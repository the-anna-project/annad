package main

import (
	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/spec"
)

type mainO struct{}

// GetObjectID returns the hardcoded object ID 56139b39e2f979be.
func (m mainO) GetObjectID() spec.ObjectID {
	return spec.ObjectID("56139b39e2f979be")
}

func (m mainO) GetObjectType() spec.ObjectType {
	return common.ObjectType.Main
}

// GetState is not implemented for this object.
func (m mainO) GetState() spec.State {
	return nil
}

// SetState is not implemented for this object.
func (m mainO) SetState(state spec.State) {
}
