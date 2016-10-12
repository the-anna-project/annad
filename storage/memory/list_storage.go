package memory

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *storage) PopFromList(key string) (string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call PopFromList")

	result, err := s.RedisStorage.PopFromList(key)
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *storage) PushToList(key string, element string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call PushToList")

	err := s.RedisStorage.PushToList(key, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
