package redis

import (
	"github.com/cenk/backoff"
	"github.com/garyburd/redigo/redis"

	"github.com/xh3b4sd/anna/spec"
)

// TODO test
func (s *storage) PopFromList(key string) (string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call PopFromList")

	var result string
	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		var err error
		result, err = redis.String(conn.Do("BRPOP", s.withPrefix(key)))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("PopFromList", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

// TODO test
func (s *storage) PushToList(key string, element string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call PushToList")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("LPUSH", s.withPrefix(key), element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("PushToList", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
