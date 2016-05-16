package strategy

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *strategy) GetID() spec.ObjectID {
	return s.ID
}

func (s *strategy) GetType() spec.ObjectType {
	return s.Type
}
