package memory

func (s *storage) GetAllFromSet(key string) ([]string, error) {
	s.Service().Log().Line("func", "GetAllFromSet")

	result, err := s.RedisStorage.GetAllFromSet(key)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *storage) PushToSet(key string, element string) error {
	s.Service().Log().Line("func", "PushToSet")

	err := s.RedisStorage.PushToSet(key, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) RemoveFromSet(key string, element string) error {
	s.Service().Log().Line("func", "RemoveFromSet")

	err := s.RedisStorage.RemoveFromSet(key, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) WalkSet(key string, closer <-chan struct{}, cb func(element string) error) error {
	s.Service().Log().Line("func", "WalkSet")

	err := s.RedisStorage.WalkSet(key, closer, cb)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
