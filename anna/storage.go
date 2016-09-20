package main

import (
	"github.com/cenk/backoff"

	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
	"github.com/xh3b4sd/anna/storage/memory"
	"github.com/xh3b4sd/anna/storage/redis"
)

func createStorageCollection(newLog spec.Log, flags Flags) (spec.StorageCollection, error) {
	newFeatureStorage, err := createFeatureStorage(newLog, flags)
	if err != nil {
		return nil, maskAny(err)
	}
	newGeneralStorage, err := createGeneralStorage(newLog, flags)
	if err != nil {
		return nil, maskAny(err)
	}

	newCollectionConfig := storage.DefaultCollectionConfig()
	newCollectionConfig.FeatureStorage = newFeatureStorage
	newCollectionConfig.GeneralStorage = newGeneralStorage
	newCollection, err := storage.NewCollection(newCollectionConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newCollection, nil
}

func createFeatureStorage(newLog spec.Log, flags Flags) (spec.Storage, error) {
	var newStorage spec.Storage
	var err error

	switch flags.Storage {
	case "redis":
		// dial
		newDialConfig := redis.DefaultDialConfig()
		newDialConfig.Addr = flags.RedisFeatureStorageAddr
		// pool
		newPoolConfig := redis.DefaultPoolConfig()
		newPoolConfig.Dial = redis.NewDial(newDialConfig)
		// storage
		newStorageConfig := redis.DefaultStorageConfig()
		newStorageConfig.BackOffFactory = func() spec.BackOff {
			return backoff.NewExponentialBackOff()
		}
		newStorageConfig.Log = newLog
		newStorageConfig.Instrumentation, err = createPrometheusInstrumentation([]string{"Feature", "Storage", "Redis"})
		if err != nil {
			return nil, maskAny(err)
		}
		newStorageConfig.Pool = redis.NewPool(newPoolConfig)
		newStorageConfig.Prefix = flags.RedisStoragePrefix
		newStorage, err = redis.NewStorage(newStorageConfig)
		if err != nil {
			return nil, maskAny(err)
		}
	case "memory":
		newStorage, err = memory.NewStorage(memory.DefaultStorageConfig())
		if err != nil {
			return nil, maskAny(err)
		}
	default:
		return nil, maskAnyf(invalidStorageFlagError, "%s", flags.Storage)
	}

	return newStorage, nil
}

func createGeneralStorage(newLog spec.Log, flags Flags) (spec.Storage, error) {
	var newStorage spec.Storage
	var err error

	switch flags.Storage {
	case "redis":
		// dial
		newDialConfig := redis.DefaultDialConfig()
		newDialConfig.Addr = flags.RedisGeneralStorageAddr
		// pool
		newPoolConfig := redis.DefaultPoolConfig()
		newPoolConfig.Dial = redis.NewDial(newDialConfig)
		// storage
		newStorageConfig := redis.DefaultStorageConfig()
		newStorageConfig.BackOffFactory = func() spec.BackOff {
			return backoff.NewExponentialBackOff()
		}
		newStorageConfig.Log = newLog
		newStorageConfig.Instrumentation, err = createPrometheusInstrumentation([]string{"General", "Storage", "Redis"})
		if err != nil {
			return nil, maskAny(err)
		}
		newStorageConfig.Pool = redis.NewPool(newPoolConfig)
		newStorageConfig.Prefix = flags.RedisStoragePrefix
		newStorage, err = redis.NewStorage(newStorageConfig)
		if err != nil {
			return nil, maskAny(err)
		}
	case "memory":
		newStorage, err = memory.NewStorage(memory.DefaultStorageConfig())
		if err != nil {
			return nil, maskAny(err)
		}
	default:
		return nil, maskAnyf(invalidStorageFlagError, "%s", flags.Storage)
	}

	return newStorage, nil
}
