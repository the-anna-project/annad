package memory

import (
	"sync"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/cenk/backoff"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/redis"
)

const (
	// ObjectType represents the object type of the memory storage object. This
	// is used e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "memory-storage"
)

// StorageConfig represents the configuration used to create a new memory
// storage object.
type StorageConfig struct {
	// Dependencies.

	Log spec.Log
}

// DefaultStorageConfig provides a default configuration to create a new memory
// storage object by best effort.
func DefaultStorageConfig() StorageConfig {
	newConfig := StorageConfig{
		// Dependencies.
		Log: log.New(log.DefaultConfig()),
	}

	return newConfig
}

// NewStorage creates a new configured memory storage object. Therefore it
// manages an in-memory redis instance which can be shut down using the
// configured closer. This is used for local development.
func NewStorage(config StorageConfig) (spec.Storage, error) {
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
	newRedisStorageConfig.BackOffFactory = func() spec.BackOff {
		return backoff.NewExponentialBackOff()
	}
	newRedisStorage, err := redis.NewStorage(newRedisStorageConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	newStorage := &storage{
		StorageConfig: config,

		Closer:       closer,
		ID:           id.MustNew(),
		RedisStorage: newRedisStorage,
		ShutdownOnce: sync.Once{},
		Type:         ObjectType,
	}

	return newStorage, nil
}

// MustNew creates either a new default configured storage object, or panics.
func MustNew() spec.Storage {
	newStorage, err := NewStorage(DefaultStorageConfig())
	if err != nil {
		panic(err)
	}

	return newStorage
}

type storage struct {
	StorageConfig

	Closer       chan struct{}
	ID           spec.ObjectID
	RedisStorage spec.Storage
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (s *storage) GetRandom() (string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call GetRandom")

	result, err := s.RedisStorage.GetRandom()
	if err != nil {
		return "", maskAny(err)
	}

	return result, nil
}

func (s *storage) GetStringMap(key string) (map[string]string, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call GetStringMap")

	result, err := s.RedisStorage.GetStringMap(key)
	if err != nil {
		return nil, maskAny(err)
	}

	return result, nil
}

func (s *storage) SetStringMap(key string, stringMap map[string]string) error {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call SetStringMap")

	err := s.RedisStorage.SetStringMap(key, stringMap)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

func (s *storage) Shutdown() {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call Shutdown")

	s.ShutdownOnce.Do(func() {
		close(s.Closer)
	})
}
