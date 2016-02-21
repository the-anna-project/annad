package server

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *server) MarshalJSON() ([]byte, error) {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 15}, "call MarshalJSON")
	return nil, nil
}
