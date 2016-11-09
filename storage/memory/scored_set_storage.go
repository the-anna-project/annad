package memory

func (s *storage) GetElementsByScore(key string, score float64, maxElements int) ([]string, error) {
	s.Service().Log().Line("func", "GetElementsByScore")

	result, err := s.RedisStorage.GetElementsByScore(key, score, maxElements)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *storage) GetHighestScoredElements(key string, maxElements int) ([]string, error) {
	s.Service().Log().Line("func", "GetHighestScoredElements")

	result, err := s.RedisStorage.GetHighestScoredElements(key, maxElements)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *storage) RemoveScoredElement(key string, element string) error {
	s.Service().Log().Line("func", "RemoveScoredElement")

	err := s.RedisStorage.RemoveScoredElement(key, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) SetElementByScore(key, element string, score float64) error {
	s.Service().Log().Line("func", "SetElementByScore")

	err := s.RedisStorage.SetElementByScore(key, element, score)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) WalkScoredSet(key string, closer <-chan struct{}, cb func(element string, score float64) error) error {
	s.Service().Log().Line("func", "WalkScoredSet")

	err := s.RedisStorage.WalkScoredSet(key, closer, cb)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
