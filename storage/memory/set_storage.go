package memory

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *storage) GetAllFromSet(key string) ([]string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call GetAllFromSet")

	result, err := s.RedisStorage.GetAllFromSet(key)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *storage) PushToSet(key string, element string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call PushToSet")

	err := s.RedisStorage.PushToSet(key, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) RemoveFromSet(key string, element string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call RemoveFromSet")

	err := s.RedisStorage.RemoveFromSet(key, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) WalkSet(key string, closer <-chan struct{}, cb func(element string) error) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call WalkSet")

	err := s.RedisStorage.WalkSet(key, closer, cb)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
