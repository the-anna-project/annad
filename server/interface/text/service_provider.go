package text

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *server) Service() spec.ServiceCollection {
	return s.ServiceCollection
}
