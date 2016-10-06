package redis

import (
	"strconv"
	"sync"

	"github.com/cenk/backoff"
	"github.com/garyburd/redigo/redis"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/instrumentation/memory"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectType represents the object type of the redis storage object. This is
	// used e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "redis-storage"
)

// StorageConfig represents the configuration used to create a new redis
// storage object.
type StorageConfig struct {
	// Dependencies.
	Instrumentation spec.Instrumentation
	Log             spec.Log
	Pool            *redis.Pool

	// Settings.

	// BackOffFactory is supposed to be able to create a new spec.BackOff. Retry
	// implementations can make use of this to decide when to retry.
	BackOffFactory func() spec.BackOff

	Prefix string
}

// DefaultStorageConfigWithConn provides a configuration that can be mocked
// using a redis connection. This is used for testing.
func DefaultStorageConfigWithConn(redisConn redis.Conn) StorageConfig {
	newPoolConfig := DefaultPoolConfig()
	newMockDialConfig := defaultMockDialConfig()
	newMockDialConfig.RedisConn = redisConn
	newPoolConfig.Dial = newMockDial(newMockDialConfig)
	newPool := NewPool(newPoolConfig)

	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.Pool = newPool

	return newStorageConfig
}

// DefaultStorageConfig provides a default configuration to create a new redis
// storage object by best effort.
func DefaultStorageConfig() StorageConfig {
	newInstrumentation, err := memory.NewInstrumentation(memory.DefaultInstrumentationConfig())
	if err != nil {
		panic(err)
	}

	newConfig := StorageConfig{
		// Dependencies.
		Instrumentation: newInstrumentation,
		Log:             log.New(log.DefaultConfig()),
		Pool:            NewPool(DefaultPoolConfig()),

		// Settings.
		BackOffFactory: func() spec.BackOff {
			return &backoff.StopBackOff{}
		},
		Prefix: "prefix",
	}

	return newConfig
}

// NewStorage creates a new configured redis storage object.
func NewStorage(config StorageConfig) (spec.Storage, error) {
	newStorage := &storage{
		StorageConfig: config,

		ID:    id.MustNew(),
		Mutex: sync.Mutex{},
		Type:  ObjectType,
	}

	// Dependencies.
	if newStorage.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newStorage.Pool == nil {
		return nil, maskAnyf(invalidConfigError, "connection pool must not be empty")
	}
	// Settings.
	if newStorage.BackOffFactory == nil {
		return nil, maskAnyf(invalidConfigError, "backoff factory must not be empty")
	}
	if newStorage.Prefix == "" {
		return nil, maskAnyf(invalidConfigError, "prefix must not be empty")
	}

	newStorage.Log.Register(newStorage.GetType())

	return newStorage, nil
}

type storage struct {
	StorageConfig

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (s *storage) Get(key string) (string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call Get")

	var result string
	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		var err error
		result, err = redis.String(conn.Do("GET", s.withPrefix(key)))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("Get", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *storage) GetElementsByScore(key string, score float64, maxElements int) ([]string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call GetElementsByScore")

	var result []string
	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		values, err := redis.Values(conn.Do("ZREVRANGEBYSCORE", s.withPrefix(key), score, score, "LIMIT", 0, maxElements))
		if err != nil {
			return maskAny(err)
		}

		for _, v := range values {
			result = append(result, v.(string))
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("GetElementsByScore", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *storage) GetHighestScoredElements(key string, maxElements int) ([]string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call GetHighestScoredElements")

	var result []string
	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		values, err := redis.Values(conn.Do("ZREVRANGE", s.withPrefix(key), 0, maxElements-1, "WITHSCORES"))
		if err != nil {
			return maskAny(err)
		}

		for _, v := range values {
			result = append(result, v.(string))
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("GetHighestScoredElements", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *storage) GetRandomKey() (string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call GetRandomKey")

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

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("GetRandomKey", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *storage) GetStringMap(key string) (map[string]string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call GetStringMap")

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

	err = backoff.RetryNotify(s.Instrumentation.WrapFunc("GetStringMap", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

// TODO
func (s *storage) PopFromList(key string) (string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call PushToSet")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("SADD", s.withPrefix(key)))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("PushToSet", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return "", maskAny(err)
	}

	return "", nil
}

// TODO
func (s *storage) PushToList(key string, element string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call PushToSet")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("SADD", s.withPrefix(key), element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("PushToSet", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) PushToSet(key string, element string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call PushToSet")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("SADD", s.withPrefix(key), element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("PushToSet", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) RemoveFromSet(key string, element string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call RemoveFromSet")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("SREM", s.withPrefix(key), element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("RemoveFromSet", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) RemoveScoredElement(key string, element string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call RemoveScoredElement")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("ZREM", s.withPrefix(key), element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("RemoveScoredElement", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
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

func (s *storage) SetElementByScore(key, element string, score float64) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call SetElementByScore")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("ZADD", s.withPrefix(key), score, element))
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("SetElementByScore", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) SetStringMap(key string, stringMap map[string]string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call SetStringMap")

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

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("SetStringMap", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) WalkScoredElements(key string, closer <-chan struct{}, cb func(element string, score float64) error) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call WalkScoredElements")

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

			reply, err := redis.Values(conn.Do("ZSCAN", s.withPrefix(key), cursor, "COUNT", 100))
			if err != nil {
				return maskAny(err)
			}

			cursor := reply[0].(int64)
			values := reply[1].([]string)

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

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("WalkScoredElements", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) WalkSet(key string, closer <-chan struct{}, cb func(element string) error) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call WalkSet")

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

			cursor := reply[0].(int64)
			values := reply[1].([]string)

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

	err := backoff.RetryNotify(s.Instrumentation.WrapFunc("WalkSet", action), s.BackOffFactory(), s.retryErrorLogger)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
