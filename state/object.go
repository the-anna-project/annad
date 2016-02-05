package state

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *state) GetObjectID() spec.ObjectID {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.ObjectID
}

func (s *state) GetObjectType() spec.ObjectType {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.ObjectType
}

// GetState is not implemented for this object.
func (s *state) GetState() spec.State {
	return nil
}

// SetState is not implemented for this object.
func (s *state) SetState(state spec.State) {
}
