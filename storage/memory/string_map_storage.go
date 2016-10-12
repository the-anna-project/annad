package memory

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *storage) GetStringMap(key string) (map[string]string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call GetStringMap")

	result, err := s.RedisStorage.GetStringMap(key)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *storage) SetStringMap(key string, stringMap map[string]string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call SetStringMap")

	err := s.RedisStorage.SetStringMap(key, stringMap)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
