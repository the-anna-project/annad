package random

import (
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

func (s *service) Service() servicespec.Collection {
	return s.ServiceCollection
}
