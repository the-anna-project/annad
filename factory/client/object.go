package factoryclient

import (
	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/spec"
)

// GetObjectID returns the hardcoded object ID 1501816570b0e088.
func (c *client) GetObjectID() spec.ObjectID {
	return spec.ObjectID("1501816570b0e088")
}

func (c *client) GetObjectType() spec.ObjectType {
	return common.ObjectType.FactoryClient
}

// GetState is not implemented for this object.
func (c *client) GetState() spec.State {
	return nil
}

// SetState is not implemented for this object.
func (c *client) SetState(state spec.State) {
}
