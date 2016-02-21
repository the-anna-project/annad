package server

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *server) GetID() spec.ObjectID {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.ID
}

func (s *server) GetType() spec.ObjectType {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.Type
}
