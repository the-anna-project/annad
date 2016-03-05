package storage

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	ObjectTypeMemoryStorage spec.ObjectType = "memory-storage"
)

type MemoryStorageConfig struct {
	Log     spec.Log
	Storage spec.Storage
}

func DefaultMemoryStorageConfig() MemoryStorageConfig {
	newPoolConfig := DefaultRedisPoolConfig()
	newPoolConfig.Dial = NewMemoryDial(DefaultMemoryDialConfig())
	newPool := NewRedisPool(newPoolConfig)

	newStorageConfig := DefaultRedisStorageConfig()
	newStorageConfig.Pool = newPool

	newConfig := MemoryStorageConfig{
		Log:     log.NewLog(log.DefaultConfig()),
		Storage: NewRedisStorage(newStorageConfig),
	}

	return newConfig
}

func NewMemoryStorage(config MemoryStorageConfig) spec.Storage {
	newStorage := &memoryStorage{
		ID:                  id.NewObjectID(id.Hex128),
		MemoryStorageConfig: config,
		Mutex:               sync.Mutex{},
		Type:                ObjectTypeMemoryStorage,
	}

	newStorage.Log.Register(newStorage.GetType())

	return newStorage
}

type memoryStorage struct {
	ID spec.ObjectID
	MemoryStorageConfig
	Mutex sync.Mutex `json:"-"`
	Type  spec.ObjectType
}

func (ms *memoryStorage) Get(key string) (string, error) {
	return ms.Storage.Get(key)
}

func (ms *memoryStorage) GetElementsByScore(key string, score float32, maxElements int) ([]string, error) {
	return ms.Storage.GetElementsByScore(key, score, maxElements)
}

func (ms *memoryStorage) GetHighestScoredElements(key string, maxElements int) ([]string, error) {
	return ms.Storage.GetHighestScoredElements(key, maxElements)
}

func (ms *memoryStorage) Set(key, value string) error {
	return ms.Storage.Set(key, value)
}
