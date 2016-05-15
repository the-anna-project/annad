package scheduler

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *scheduler) GetID() spec.ObjectID {
	return s.ID
}

func (s *scheduler) GetType() spec.ObjectType {
	return s.Type
}
