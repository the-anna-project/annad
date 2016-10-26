package memory

import "github.com/xh3b4sd/anna/spec"

func (s *storage) Get(key string) (string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call Get")

	result, err := s.RedisStorage.Get(key)
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *storage) GetRandom() (string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call GetRandom")

	result, err := s.RedisStorage.GetRandom()
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *storage) Remove(key string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call Remove")

	err := s.RedisStorage.Remove(key)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) Set(key, value string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call Set")

	err := s.RedisStorage.Set(key, value)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) WalkKeys(glob string, closer <-chan struct{}, cb func(key string) error) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call WalkKeys")

	err := s.RedisStorage.WalkKeys(glob, closer, cb)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
