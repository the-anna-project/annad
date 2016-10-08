package memory

import (
	"time"

	"github.com/alicebob/miniredis"

	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/redis"
)

const (
	// ObjectType represents the object type of the memory storage object. This
	// is used e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "memory-storage"
)

// DefaultStorageConfig provides a default configuration to create a new memory
// storage object by best effort.
func DefaultStorageConfig() redis.StorageConfig {
	newConfig := DefaultStorageConfigWithCloser(nil)
	return newConfig
}

// DefaultStorageConfigWithCloser provides a configuration that manages an
// in-memory redis instance which can be shut down using the provided closer.
// This is used for local development.
func DefaultStorageConfigWithCloser(closer chan struct{}) redis.StorageConfig {
	// miniredis
	addrChan := make(chan string, 1)
	go func() {
		s, err := miniredis.Run()
		if err != nil {
			panic(err)
		}
		addrChan <- s.Addr()

		<-closer
		s.Close()
	}()

	// dial
	newDialConfig := redis.DefaultDialConfig()
	select {
	case <-time.After(1 * time.Second):
		panic("starting miniredis timed out")
	case addr := <-addrChan:
		newDialConfig.Addr = addr
	}

	// pool
	newPoolConfig := redis.DefaultPoolConfig()
	newPoolConfig.Dial = redis.NewDial(newDialConfig)
	newPool := redis.NewPool(newPoolConfig)

	// storage
	newStorageConfig := redis.DefaultStorageConfig()
	newStorageConfig.Pool = newPool

	return newStorageConfig
}

// NewStorage creates a new configured memory storage object.
func NewStorage(config redis.StorageConfig) (spec.Storage, error) {
	newStorage, err := redis.NewStorage(config)
	if err != nil {
		return nil, maskAny(err)
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
