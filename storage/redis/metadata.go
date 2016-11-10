package redis

func (s *storage) GetMetadata() map[string]string {
	return s.Metadata
}
