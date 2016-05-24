// Package redisstorage implements spec.Storage and provides functionality to
// persist data in redis.
package redisstorage

import (
	"strconv"
	"sync"

	"github.com/cenk/backoff"
	"github.com/garyburd/redigo/redis"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/instrumentation/prometheus"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeRedisStorage represents the object type of the redis storage
	// object. This is used e.g. to register itself to the logger.
	ObjectTypeRedisStorage spec.ObjectType = "redis-storage"
)

// Config represents the configuration used to create a new redis storage
// object.
type Config struct {
	// Dependencies.
	Instrumentation spec.Instrumentation
	Log             spec.Log
	Pool            *redis.Pool

	// Settings.

	// BackOffFactory is supposed to be able to create a new spec.BackOff. Retry
	// implementations can make use of this to decide when to retry.
	BackOffFactory func() spec.BackOff
}

// DefaultConfigWithConn provides a configuration that can be mocked using a
// redis connection. This is used for testing.
func DefaultConfigWithConn(redisConn redis.Conn) Config {
	newPoolConfig := DefaultRedisPoolConfig()
	newMockDialConfig := defaultMockDialConfig()
	newMockDialConfig.RedisConn = redisConn
	newPoolConfig.Dial = newMockDial(newMockDialConfig)
	newPool := NewRedisPool(newPoolConfig)

	newStorageConfig := DefaultConfig()
	newStorageConfig.Pool = newPool

	return newStorageConfig
}

// DefaultConfig provides a default configuration to create a new redis storage
// object by best effort.
func DefaultConfig() Config {
	newPrometheusConfig := prometheus.DefaultConfig()
	newPrometheusConfig.Prefixes = append(newPrometheusConfig.Prefixes, "Storage", "Redis")
	newInstrumentation, err := prometheus.New(newPrometheusConfig)
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.
		Instrumentation: newInstrumentation,
		Log:             log.NewLog(log.DefaultConfig()),
		Pool:            NewRedisPool(DefaultRedisPoolConfig()),

		// Settings.
		BackOffFactory: func() spec.BackOff {
			return &backoff.StopBackOff{}
		},
	}

	return newConfig
}

// NewRedisStorage creates a new configured redis storage object.
func NewRedisStorage(config Config) (spec.Storage, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newStorage := &storage{
		Config: config,

		ID:    newID,
		Mutex: sync.Mutex{},
		Type:  ObjectTypeRedisStorage,
	}

	if newStorage.BackOffFactory == nil {
		return nil, maskAnyf(invalidConfigError, "backoff factory must not be empty")
	}
	if newStorage.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newStorage.Pool == nil {
		return nil, maskAnyf(invalidConfigError, "connection pool must not be empty")
	}

	newStorage.Log.Register(newStorage.GetType())

	return newStorage, nil
}

type storage struct {
	Config

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (s *storage) Get(key string) (string, error) {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call Get")

	var result string
	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		var err error
		result, err = redis.String(conn.Do("GET", key))
		if !IsNil(err) {
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
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call GetElementsByScore")

	var result []string
	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		values, err := redis.Values(conn.Do("ZREVRANGEBYSCORE", key, score, score, "LIMIT", 0, maxElements))
		if !IsNil(err) {
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

func (s *storage) GetStringMap(key string) (map[string]string, error) {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call GetStringMap")

	var result map[string]string
	var err error
	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		result, err = redis.StringMap(conn.Do("HGETALL", key))
		if !IsNil(err) {
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

func (s *storage) GetHighestScoredElements(key string, maxElements int) ([]string, error) {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call GetHighestScoredElements")

	var result []string
	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		values, err := redis.Values(conn.Do("ZREVRANGE", key, 0, maxElements-1, "WITHSCORES"))
		if !IsNil(err) {
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

func (s *storage) PushToSet(key string, element string) error {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call PushToSet")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("SADD", key, element))
		if !IsNil(err) {
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
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call RemoveFromSet")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("SREM", key, element))
		if !IsNil(err) {
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
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call RemoveScoredElement")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("ZREM", key, element))
		if !IsNil(err) {
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
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call Set")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		reply, err := redis.String(conn.Do("SET", key, value))
		if !IsNil(err) {
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
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call SetElementByScore")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		_, err := redis.Int(conn.Do("ZADD", key, score, element))
		if !IsNil(err) {
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
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call SetStringMap")

	action := func() error {
		conn := s.Pool.Get()
		defer conn.Close()

		reply, err := redis.String(conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(stringMap)...))
		if !IsNil(err) {
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
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call WalkScoredElements")

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

			reply, err := redis.Values(conn.Do("ZSCAN", key, cursor, "COUNT", 100))
			if !IsNil(err) {
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
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call WalkSet")

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

			reply, err := redis.Values(conn.Do("SSCAN", key, cursor, "COUNT", 100))
			if !IsNil(err) {
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
