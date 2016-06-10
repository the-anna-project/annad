package strategy

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *static) GetID() spec.ObjectID {
	return s.ID
}

func (s *static) GetType() spec.ObjectType {
	return s.Type
}
