package redis

import (
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

func (s *storage) Service() servicespec.Collection {
	return s.ServiceCollection
}
