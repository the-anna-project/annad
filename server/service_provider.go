package server

import (
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

func (s *server) Service() servicespec.Collection {
	return s.ServiceCollection
}
