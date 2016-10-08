package main

import (
	"github.com/cenk/backoff"

	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
	"github.com/xh3b4sd/anna/storage/memory"
	"github.com/xh3b4sd/anna/storage/redis"
)

func newStorageCollection(newLog spec.Log, closer chan struct{}, flags Flags) (spec.StorageCollection, error) {
	newFeatureStorage, err := newConfiguredStorage(newLog, closer, flags.Storage, flags.RedisStoragePrefix, flags.RedisFeatureStorageAddr)
	if err != nil {
		return nil, maskAny(err)
	}
	newGeneralStorage, err := newConfiguredStorage(newLog, closer, flags.Storage, flags.RedisStoragePrefix, flags.RedisGeneralStorageAddr)
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

func newConfiguredStorage(newLog spec.Log, closer chan struct{}, storageType, storagePrefix, storageAddr string) (spec.Storage, error) {
	var newStorageConfig redis.StorageConfig
	var err error

	switch storageType {
	case "redis":
		newStorageConfig = redis.DefaultStorageConfigWithAddr(storageAddr)
		newStorageConfig.BackOffFactory = func() spec.BackOff {
			return backoff.NewExponentialBackOff()
		}
		newStorageConfig.Instrumentation, err = newPrometheusInstrumentation([]string{"Feature", "Storage", "Redis"})
		if err != nil {
			return nil, maskAny(err)
		}
		newStorageConfig.Log = newLog
		newStorageConfig.Prefix = storagePrefix
	case "memory":
		// storage
		newStorageConfig = memory.DefaultStorageConfigWithCloser(closer)
	default:
		return nil, maskAnyf(invalidStorageFlagError, "%s", storageType)
	}

	newStorage, err := redis.NewStorage(newStorageConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newStorage, nil
}
