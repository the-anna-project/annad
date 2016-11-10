package memory

func (s *storage) Get(key string) (string, error) {
	s.Service().Log().Line("func", "Get")

	result, err := s.RedisStorage.Get(key)
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *storage) GetRandom() (string, error) {
	s.Service().Log().Line("func", "GetRandom")

	result, err := s.RedisStorage.GetRandom()
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *storage) Remove(key string) error {
	s.Service().Log().Line("func", "Remove")

	err := s.RedisStorage.Remove(key)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) Set(key, value string) error {
	s.Service().Log().Line("func", "Set")

	err := s.RedisStorage.Set(key, value)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) WalkKeys(glob string, closer <-chan struct{}, cb func(key string) error) error {
	s.Service().Log().Line("func", "WalkKeys")

	err := s.RedisStorage.WalkKeys(glob, closer, cb)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
