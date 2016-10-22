package memory

import (
	"github.com/xh3b4sd/anna/spec"
)

func (s *storage) GetElementsByScore(key string, score float64, maxElements int) ([]string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call GetElementsByScore")

	result, err := s.RedisStorage.GetElementsByScore(key, score, maxElements)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *storage) GetHighestScoredElements(key string, maxElements int) ([]string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call GetHighestScoredElements")

	result, err := s.RedisStorage.GetHighestScoredElements(key, maxElements)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *storage) RemoveScoredElement(key string, element string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call RemoveScoredElement")

	err := s.RedisStorage.RemoveScoredElement(key, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) SetElementByScore(key, element string, score float64) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call SetElementByScore")

	err := s.RedisStorage.SetElementByScore(key, element, score)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) WalkScoredSet(key string, closer <-chan struct{}, cb func(element string, score float64) error) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call WalkScoredSet")

	err := s.RedisStorage.WalkScoredSet(key, closer, cb)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
