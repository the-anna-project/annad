package main

import (
	"github.com/cenk/backoff"

	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
	"github.com/xh3b4sd/anna/storage/redis"
)

func createStorage(newLog spec.Log) (spec.Storage, error) {
	var newStorage spec.Storage
	var err error

	switch globalFlags.Storage {
	case "redis":
		// dial
		newDialConfig := redis.DefaultDialConfig()
		newDialConfig.Addr = globalFlags.StorageAddr
		// pool
		newPoolConfig := redis.DefaultPoolConfig()
		newPoolConfig.Dial = redis.NewDial(newDialConfig)
		// storage
		newStorageConfig := redis.DefaultStorageConfig()
		newStorageConfig.BackOffFactory = func() spec.BackOff {
			return backoff.NewExponentialBackOff()
		}
		newStorageConfig.Log = newLog
		newStorageConfig.Instrumentation, err = createPrometheusInstrumentation([]string{"Storage", "Redis"})
		if err != nil {
			return nil, maskAny(err)
		}
		newStorageConfig.Pool = redis.NewPool(newPoolConfig)
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
		return nil, maskAnyf(invalidStorageFlagError, "%s", globalFlags.Storage)
	}

	return newStorage, nil
}
