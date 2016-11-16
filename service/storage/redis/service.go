package redis

import (
	"strconv"
	"sync"

	"github.com/cenk/backoff"
	"github.com/garyburd/redigo/redis"

	objectspec "github.com/the-anna-project/spec/object"
	servicespec "github.com/the-anna-project/spec/service"
)

// New creates a new redis storage service.
func New() servicespec.StorageService {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.ServiceCollection

	// Settings.

	address string
	// backoffFactory is supposed to be able to create a new spec.Backoff. Retry
	// implementations can make use of this to decide when to retry.
	backoffFactory func() objectspec.Backoff
	metadata       map[string]string
	pool           *redis.Pool
	prefix         string
	shutdownOnce   sync.Once
}

func (s *service) Boot() {
	id, err := s.Service().ID().New()
	if err != nil {
		panic(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"kind": "redis",
		"name": "storage",
		"type": "service",
	}

	s.backoffFactory = func() objectspec.Backoff {
		return &backoff.StopBackOff{}
	}
	// pool
	newDialConfig := DefaultDialConfig()
	newDialConfig.Addr = s.address
	newPoolConfig := DefaultPoolConfig()
	newPoolConfig.Dial = NewDial(newDialConfig)
	s.pool = NewPool(newPoolConfig)
	s.prefix = "prefix"
	s.shutdownOnce = sync.Once{}
}

func (s *service) Get(key string) (string, error) {
	s.Service().Log().Line("func", "Get")

	errors := make(chan error, 1)

	var result string
	action := func() error {
		conn := s.pool.Get()
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

	err := backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("Get", action), s.backoffFactory(), s.retryErrorLogger)
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

func (s *service) GetAllFromSet(key string) ([]string, error) {
	s.Service().Log().Line("func", "GetAllFromSet")

	var result []string
	action := func() error {
		conn := s.pool.Get()
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

	err := backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("GetAllFromSet", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *service) GetElementsByScore(key string, score float64, maxElements int) ([]string, error) {
	s.Service().Log().Line("func", "GetElementsByScore")

	var result []string
	var err error
	action := func() error {
		conn := s.pool.Get()
		defer conn.Close()

		result, err = redis.Strings(conn.Do("ZREVRANGEBYSCORE", s.withPrefix(key), score, score, "LIMIT", 0, maxElements))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err = backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("GetElementsByScore", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *service) GetHighestScoredElements(key string, maxElements int) ([]string, error) {
	s.Service().Log().Line("func", "GetHighestScoredElements")

	var result []string
	var err error
	action := func() error {
		conn := s.pool.Get()
		defer conn.Close()

		result, err = redis.Strings(conn.Do("ZREVRANGE", s.withPrefix(key), 0, maxElements, "WITHSCORES"))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err = backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("GetHighestScoredElements", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *service) GetRandom() (string, error) {
	s.Service().Log().Line("func", "GetRandom")

	var result string
	action := func() error {
		conn := s.pool.Get()
		defer conn.Close()

		var err error
		result, err = redis.String(conn.Do("RANDOMKEY"))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("GetRandom", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *service) GetStringMap(key string) (map[string]string, error) {
	s.Service().Log().Line("func", "GetStringMap")

	var result map[string]string
	var err error
	action := func() error {
		conn := s.pool.Get()
		defer conn.Close()

		result, err = redis.StringMap(conn.Do("HGETALL", s.withPrefix(key)))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err = backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("GetStringMap", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) PopFromList(key string) (string, error) {
	s.Service().Log().Line("func", "PopFromList")

	var result string
	action := func() error {
		conn := s.pool.Get()
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

	err := backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("PopFromList", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *service) PushToList(key string, element string) error {
	s.Service().Log().Line("func", "PushToList")

	action := func() error {
		conn := s.pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("LPUSH", s.withPrefix(key), element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("PushToList", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) PushToSet(key string, element string) error {
	s.Service().Log().Line("func", "PushToSet")

	action := func() error {
		conn := s.pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("SADD", s.withPrefix(key), element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("PushToSet", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Remove(key string) error {
	s.Service().Log().Line("func", "Remove")

	action := func() error {
		conn := s.pool.Get()
		defer conn.Close()

		_, err := redis.Int64(conn.Do("DEL", s.withPrefix(key)))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.Retry(s.Service().Instrumentor().WrapFunc("Remove", action), s.backoffFactory())
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) RemoveFromSet(key string, element string) error {
	s.Service().Log().Line("func", "RemoveFromSet")

	action := func() error {
		conn := s.pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("SREM", s.withPrefix(key), element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("RemoveFromSet", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) RemoveScoredElement(key string, element string) error {
	s.Service().Log().Line("func", "RemoveScoredElement")

	action := func() error {
		conn := s.pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("ZREM", s.withPrefix(key), element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("RemoveScoredElement", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Service() servicespec.ServiceCollection {
	return s.serviceCollection
}

func (s *service) Set(key, value string) error {
	s.Service().Log().Line("func", "Set")

	action := func() error {
		conn := s.pool.Get()
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

	err := backoff.Retry(s.Service().Instrumentor().WrapFunc("Set", action), s.backoffFactory())
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) SetAddress(a string) {
	s.address = a
}

func (s *service) SetBackoffFactory(bf func() objectspec.Backoff) {
	s.backoffFactory = bf
}

func (s *service) SetElementByScore(key, element string, score float64) error {
	s.Service().Log().Line("func", "SetElementByScore")

	action := func() error {
		conn := s.pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("ZADD", s.withPrefix(key), score, element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("SetElementByScore", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) SetPrefix(p string) {
	s.prefix = p
}

func (s *service) SetServiceCollection(sc servicespec.ServiceCollection) {
	s.serviceCollection = sc
}

func (s *service) SetStringMap(key string, stringMap map[string]string) error {
	s.Service().Log().Line("func", "SetStringMap")

	action := func() error {
		conn := s.pool.Get()
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

	err := backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("SetStringMap", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) Shutdown() {
	s.Service().Log().Line("func", "Shutdown")

	s.shutdownOnce.Do(func() {
		s.pool.Close()
	})
}

func (s *service) WalkKeys(glob string, closer <-chan struct{}, cb func(key string) error) error {
	s.Service().Log().Line("func", "WalkKeys")

	action := func() error {
		conn := s.pool.Get()
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

			reply, err := redis.Values(conn.Do("SCAN", cursor, "MATCH", glob, "COUNT", 100))
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

	err := backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("WalkKeys", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) WalkScoredSet(key string, closer <-chan struct{}, cb func(element string, score float64) error) error {
	s.Service().Log().Line("func", "WalkScoredSet")

	action := func() error {
		conn := s.pool.Get()
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

			reply, err := redis.Values(conn.Do("ZSCAN", s.withPrefix(key), cursor, "COUNT", 100))
			if err != nil {
				return maskAny(err)
			}

			cursor, values, err := parseMultiBulkReply(reply)
			if err != nil {
				return maskAny(err)
			}

			for i := range values {
				select {
				case <-closer:
					return nil
				default:
				}

				if i%2 != 0 {
					continue
				}

				score, err := strconv.ParseFloat(values[i+1], 64)
				if err != nil {
					return maskAny(err)
				}
				err = cb(values[i], score)
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

	err := backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("WalkScoredSet", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *service) WalkSet(key string, closer <-chan struct{}, cb func(element string) error) error {
	s.Service().Log().Line("func", "WalkSet")

	action := func() error {
		conn := s.pool.Get()
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

	err := backoff.RetryNotify(s.Service().Instrumentor().WrapFunc("WalkSet", action), s.backoffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
