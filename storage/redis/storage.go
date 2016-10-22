package redis

import (
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
	// pool
	newPoolConfig := DefaultPoolConfig()
	newMockDialConfig := defaultMockDialConfig()
	newMockDialConfig.RedisConn = redisConn
	newPoolConfig.Dial = newMockDial(newMockDialConfig)
	newPool := NewPool(newPoolConfig)

	// storage
	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.Pool = newPool

	return newStorageConfig
}

// DefaultStorageConfigWithAddr provides a configuration to make a redis client
// connect to the provided address. This is used for production.
func DefaultStorageConfigWithAddr(addr string) StorageConfig {
	// dial
	newDialConfig := DefaultDialConfig()
	newDialConfig.Addr = addr
	// pool
	newPoolConfig := DefaultPoolConfig()
	newPoolConfig.Dial = NewDial(newDialConfig)
	// storage
	newStorageConfig := DefaultStorageConfig()
	newStorageConfig.Pool = NewPool(newPoolConfig)

	return newStorageConfig
}

// DefaultStorageConfig provides a default configuration to create a new redis
// storage object by best effort.
func DefaultStorageConfig() StorageConfig {
	newInstrumentation, err := memory.NewInstrumentation(memory.DefaultInstrumentationConfig())
	if err != nil {
		panic(err)
	}

	newStorageConfig := StorageConfig{
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

	return newStorageConfig
}

// NewStorage creates a new configured redis storage object.
func NewStorage(config StorageConfig) (spec.Storage, error) {
	newStorage := &storage{
		StorageConfig: config,

		ID:           id.MustNew(),
		ShutdownOnce: sync.Once{},
		Type:         ObjectType,
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

	ID           spec.ObjectID
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (s *storage) Shutdown() {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call Shutdown")

	s.ShutdownOnce.Do(func() {
		s.Pool.Close()
	})
}
