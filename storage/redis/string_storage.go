package redis

import (
	"github.com/cenk/backoff"
	"github.com/garyburd/redigo/redis"

	"github.com/xh3b4sd/anna/spec"
)

func (s *storage) Get(key string) (string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call Get")

	errors := make(chan error, 1)

	var result string
	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		var err error
		result, err = redis.String(conn.Do("GET", s.withPrefix(key)))
		if IsNotFound(err) {
			// To return the not found error we need to break through the retrier.
			// Therefore we do not return the not found error here, but dispatch it to
			// the calling goroutine. Further we simply fall through and return nil to
			// finally stop the retrier.
			errors <- maskAny(err)
			return nil
		} else if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("Get", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return "", maskAny(err)
	}

	select {
	case err := <-errors:
		if err != nil {
			return "", maskAny(err)
		}
	default:
		// If there is no error, we simply fall through to return the result.
	}

	return result, nil
}

func (s *storage) GetRandom() (string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call GetRandom")

	var result string
	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		var err error
		result, err = redis.String(conn.Do("RANDOMKEY"))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("GetRandom", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *storage) Set(key, value string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call Set")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		reply, err := redis.String(conn.Do("SET", s.withPrefix(key), value))
		if err != nil {
			return maskAny(err)
		}

		if reply != "OK" {
			return maskAnyf(queryExecutionFailedError, "SET not executed correctly")
		}

		return nil
	}

	err := backoff.Retry(s.Instrumentation.WrapFunc("Set", action), s.BackOffFactory())
	if err != nil {
		return maskAny(err)
	}

	return nil
}
