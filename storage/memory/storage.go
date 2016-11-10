package memory

import (
	"sync"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/cenk/backoff"

	servicespec "github.com/xh3b4sd/anna/service/spec"
	"github.com/xh3b4sd/anna/storage/redis"
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

// StorageConfig represents the configuration used to create a new memory
// storage object.
type StorageConfig struct {
	// Dependencies.
	ServiceCollection servicespec.Collection
}

// DefaultStorageConfig provides a default configuration to create a new memory
// storage object by best effort.
func DefaultStorageConfig() StorageConfig {
	newConfig := StorageConfig{
		// Dependencies.
		ServiceCollection: nil,
	}

	return newConfig
}

// NewStorage creates a new configured memory storage object. Therefore it
// manages an in-memory redis instance which can be shut down using the
// configured closer. This is used for local development.
func NewStorage(config StorageConfig) (storagespec.Storage, error) {
	addrChan := make(chan string, 1)
	closer := make(chan struct{}, 1)
	redisAddr := ""

	go func() {
		s, err := miniredis.Run()
		if err != nil {
			panic(err)
		}
		addrChan <- s.Addr()

		<-closer
		s.Close()
	}()

	select {
	case <-time.After(1 * time.Second):
		panic("starting miniredis timed out")
	case addr := <-addrChan:
		redisAddr = addr
	}

	newRedisStorageConfig := redis.DefaultStorageConfigWithAddr(redisAddr)
	newRedisStorageConfig.BackoffFactory = func() servicespec.Backoff {
		return backoff.NewExponentialBackOff()
	}
	newRedisStorage, err := redis.NewStorage(newRedisStorageConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	newStorage := &storage{
		StorageConfig: config,

		Closer:       closer,
		RedisStorage: newRedisStorage,
		ShutdownOnce: sync.Once{},
	}

	// Dependencies
	if newStorage.ServiceCollection == nil {
		return nil, maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	id, err := newStorage.Service().ID().New()
	if err != nil {
		return nil, maskAny(err)
	}
	newStorage.Metadata["id"] = id
	newStorage.Metadata["kind"] = "memory"
	newStorage.Metadata["name"] = "storage"
	newStorage.Metadata["type"] = "service"

	return newStorage, nil
}

// MustNew creates either a new default configured storage object, or panics.
func MustNew() storagespec.Storage {
	newStorage, err := NewStorage(DefaultStorageConfig())
	if err != nil {
		panic(err)
	}

	return newStorage
}

type storage struct {
	StorageConfig

	Closer       chan struct{}
	Metadata     map[string]string
	RedisStorage storagespec.Storage
	ShutdownOnce sync.Once
}

func (s *storage) Shutdown() {
	s.Service().Log().Line("func", "Shutdown")

	s.ShutdownOnce.Do(func() {
		close(s.Closer)
	})
}
