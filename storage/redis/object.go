package redis

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *storage) GetID() string {
	return s.ID
}

func (s *storage) GetType() spec.ObjectType {
	return s.Type
}
