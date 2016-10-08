package memory

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *storage) GetID() spec.ObjectID {
	return s.ID
}

func (s *storage) GetType() spec.ObjectType {
	return s.Type
}
