package redis

import (
	"github.com/cenk/backoff"
	"github.com/garyburd/redigo/redis"
)

func (s *storage) PopFromList(key string) (string, error) {
	s.Service().Log().Line("func", "PopFromList")

	var result string
	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		var err error
		strings, err := redis.Strings(conn.Do("BRPOP", s.withPrefix(key), 0))
		if err != nil {
			return maskAny(err)
		}
		if len(strings) != 2 {
			return maskAnyf(queryExecutionFailedError, "two elements must be returned")
		}
		result = strings[1]

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("PopFromList", action), s.BackoffFactory(), s.retryErrorLogger)
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *storage) PushToList(key string, element string) error {
	s.Service().Log().Line("func", "PushToList")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("LPUSH", s.withPrefix(key), element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("PushToList", action), s.BackoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
