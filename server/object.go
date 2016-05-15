package server

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *server) GetID() spec.ObjectID {
	return s.ID
}

func (s *server) GetType() spec.ObjectType {
	return s.Type
}
