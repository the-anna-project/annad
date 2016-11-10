package redis

import (
	"sync"

	"github.com/cenk/backoff"
	"github.com/garyburd/redigo/redis"

	"github.com/xh3b4sd/anna/instrumentation/memory"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

// StorageConfig represents the configuration used to create a new redis
// storage object.
type StorageConfig struct {
	// Dependencies.
	Instrumentation   systemspec.Instrumentation
	Pool              *redis.Pool
	ServiceCollection servicespec.Collection

	// Settings.

	// BackoffFactory is supposed to be able to create a new spec.Backoff. Retry
	// implementations can make use of this to decide when to retry.
	BackoffFactory func() servicespec.Backoff

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
		Instrumentation:   newInstrumentation,
		Pool:              NewPool(DefaultPoolConfig()),
		ServiceCollection: nil,

		// Settings.
		BackoffFactory: func() servicespec.Backoff {
			return &backoff.StopBackOff{}
		},
		Prefix: "prefix",
	}

	return newStorageConfig
}

// NewStorage creates a new configured redis storage object.
func NewStorage(config StorageConfig) (storagespec.Storage, error) {
	newStorage := &storage{
		StorageConfig: config,

		ShutdownOnce: sync.Once{},
	}

	// Dependencies.
	if newStorage.Pool == nil {
		return nil, maskAnyf(invalidConfigError, "connection pool must not be empty")
	}
	if newStorage.ServiceCollection == nil {
		return nil, maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	// Settings.
	if newStorage.BackoffFactory == nil {
		return nil, maskAnyf(invalidConfigError, "backoff factory must not be empty")
	}
	if newStorage.Prefix == "" {
		return nil, maskAnyf(invalidConfigError, "prefix must not be empty")
	}

	id, err := newStorage.Service().ID().New()
	if err != nil {
		return nil, maskAny(err)
	}
	newStorage.Metadata["id"] = id
	newStorage.Metadata["kind"] = "redis"
	newStorage.Metadata["name"] = "storage"
	newStorage.Metadata["type"] = "service"

	return newStorage, nil
}

type storage struct {
	StorageConfig

	Metadata     map[string]string
	ShutdownOnce sync.Once
}

func (s *storage) Shutdown() {
	s.Service().Log().Line("func", "Shutdown")

	s.ShutdownOnce.Do(func() {
		s.Pool.Close()
	})
}
