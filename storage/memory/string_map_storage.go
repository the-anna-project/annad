package memory

func (s *storage) GetStringMap(key string) (map[string]string, error) {
	s.Service().Log().Line("func", "GetStringMap")

	result, err := s.RedisStorage.GetStringMap(key)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *storage) SetStringMap(key string, stringMap map[string]string) error {
	s.Service().Log().Line("func", "SetStringMap")

	err := s.RedisStorage.SetStringMap(key, stringMap)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
