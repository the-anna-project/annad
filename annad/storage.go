package main

import (
	"github.com/cenk/backoff"

	servicespec "github.com/xh3b4sd/anna/service/spec"
	"github.com/xh3b4sd/anna/storage"
	"github.com/xh3b4sd/anna/storage/memory"
	"github.com/xh3b4sd/anna/storage/redis"
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

// TODO move storages to service collection
func newStorageCollection(flags Flags) (storagespec.Collection, error) {
	newFeatureStorage, err := newConfiguredStorage(flags.Storage, flags.RedisStoragePrefix, flags.RedisFeatureStorageAddr)
	if err != nil {
		return nil, maskAny(err)
	}
	newGeneralStorage, err := newConfiguredStorage(flags.Storage, flags.RedisStoragePrefix, flags.RedisGeneralStorageAddr)
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

func newConfiguredStorage(storageType, storagePrefix, storageAddr string) (storagespec.Storage, error) {
	var newStorage storagespec.Storage
	var err error

	switch storageType {
	case "redis":
		newStorageConfig := redis.DefaultStorageConfigWithAddr(storageAddr)
		newStorageConfig.BackoffFactory = func() servicespec.Backoff {
			return backoff.NewExponentialBackOff()
		}
		newStorageConfig.Instrumentation, err = newPrometheusInstrumentation([]string{"Feature", "Storage", "Redis"})
		if err != nil {
			return nil, maskAny(err)
		}
		newStorageConfig.Prefix = storagePrefix
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
		return nil, maskAnyf(invalidStorageFlagError, "%s", storageType)
	}

	return newStorage, nil
}
