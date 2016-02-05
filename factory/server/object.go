package factoryserver

import (
	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/spec"
)

// GetObjectID returns the hardcoded object ID 38a7074ed6258ee1.
func (s *server) GetObjectID() spec.ObjectID {
	return spec.ObjectID("38a7074ed6258ee1")
}

func (s *server) GetObjectType() spec.ObjectType {
	return common.ObjectType.FactoryServer
}

// GetState is not implemented for this object.
func (s *server) GetState() spec.State {
	return nil
}

// SetState is not implemented for this object.
func (s *server) SetState(state spec.State) {
}
