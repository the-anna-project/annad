package memory

func (s *storage) PopFromList(key string) (string, error) {
	s.Service().Log().Line("func", "PopFromList")

	result, err := s.RedisStorage.PopFromList(key)
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *storage) PushToList(key string, element string) error {
	s.Service().Log().Line("func", "PushToList")

	err := s.RedisStorage.PushToList(key, element)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
