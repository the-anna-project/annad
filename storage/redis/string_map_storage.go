package redis

import (
	"github.com/cenk/backoff"
	"github.com/garyburd/redigo/redis"
)

func (s *storage) GetStringMap(key string) (map[string]string, error) {
	s.Service().Log().Line("func", "GetStringMap")

	var result map[string]string
	var err error
	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		result, err = redis.StringMap(conn.Do("HGETALL", s.withPrefix(key)))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err = backoff.RetryNotify(s.Instrumentation.WrapFunc("GetStringMap", action), s.BackoffFactory(), s.retryErrorLogger)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *storage) SetStringMap(key string, stringMap map[string]string) error {
	s.Service().Log().Line("func", "SetStringMap")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		reply, err := redis.String(conn.Do("HMSET", redis.Args{}.Add(s.withPrefix(key)).AddFlat(stringMap)...))
		if err != nil {
			return maskAny(err)
		}

		if reply != "OK" {
			return maskAnyf(queryExecutionFailedError, "HMSET not executed correctly")
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("SetStringMap", action), s.BackoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
