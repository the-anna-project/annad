package scheduler

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *scheduler) GetID() spec.ObjectID {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.ID
}

func (s *scheduler) GetType() spec.ObjectType {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.Type
}
