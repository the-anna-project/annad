package redis

import (
	"github.com/cenk/backoff"
	"github.com/garyburd/redigo/redis"
)

func (s *storage) GetAllFromSet(key string) ([]string, error) {
	s.Service().Log().Line("func", "GetAllFromSet")

	var result []string
	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		values, err := redis.Values(conn.Do("SMEMBERS", s.withPrefix(key)))
		if err != nil {
			return maskAny(err)
		}

		for _, v := range values {
			result = append(result, string(v.([]uint8)))
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("GetAllFromSet", action), s.BackoffFactory(), s.retryErrorLogger)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *storage) PushToSet(key string, element string) error {
	s.Service().Log().Line("func", "PushToSet")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("SADD", s.withPrefix(key), element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("PushToSet", action), s.BackoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) RemoveFromSet(key string, element string) error {
	s.Service().Log().Line("func", "RemoveFromSet")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("SREM", s.withPrefix(key), element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("RemoveFromSet", action), s.BackoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) WalkSet(key string, closer <-chan struct{}, cb func(element string) error) error {
	s.Service().Log().Line("func", "WalkSet")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		var cursor int64

		// Start to scan the set until the cursor is 0 again. Note that we check for
		// the closer twice. At first we prevent scans in case the closer was
		// triggered directly, and second before each callback execution. That way
		// ending the walk immediately is guaranteed.
		for {
			select {
			case <-closer:
				return nil
			default:
			}

			reply, err := redis.Values(conn.Do("SSCAN", s.withPrefix(key), cursor, "COUNT", 100))
			if err != nil {
				return maskAny(err)
			}

			cursor, values, err := parseMultiBulkReply(reply)
			if err != nil {
				return maskAny(err)
			}

			for _, v := range values {
				select {
				case <-closer:
					return nil
				default:
				}

				err := cb(v)
				if err != nil {
					return maskAny(err)
				}
			}

			if cursor == 0 {
				break
			}
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("WalkSet", action), s.BackoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
