package strategy

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *strategy) GetID() spec.ObjectID {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.ID
}

func (s *strategy) GetType() spec.ObjectType {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.Type
}
