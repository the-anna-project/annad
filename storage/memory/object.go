package memorystorage

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *storage) GetID() spec.ObjectID {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.ID
}

func (s *storage) GetType() spec.ObjectType {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.Type
}
