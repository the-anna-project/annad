package mem

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *service) GetID() string {
	return s.ID
}

func (s *service) GetType() spec.ObjectType {
	return s.Type
}
