package main

import (
	"github.com/cenk/backoff"

	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
	"github.com/xh3b4sd/anna/storage/redis"
)

func (a *anna) createFeatureStorage(newLog spec.Log, prefix string) (spec.Storage, error) {
	var newStorage spec.Storage
	var err error

	switch a.Flags.Storage {
	case "redis":
		// dial
		newDialConfig := redis.DefaultDialConfig()
		newDialConfig.Addr = a.Flags.RedisFeatureStorageAddr
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
		newStorageConfig.Prefix = prefix
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
		return nil, maskAnyf(invalidStorageFlagError, "%s", a.Flags.Storage)
	}

	return newStorage, nil
}

func (a *anna) createGeneralStorage(newLog spec.Log, prefix string) (spec.Storage, error) {
	var newStorage spec.Storage
	var err error

	switch a.Flags.Storage {
	case "redis":
		// dial
		newDialConfig := redis.DefaultDialConfig()
		newDialConfig.Addr = a.Flags.RedisGeneralStorageAddr
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
		newStorageConfig.Prefix = prefix
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
		return nil, maskAnyf(invalidStorageFlagError, "%s", a.Flags.Storage)
	}

	return newStorage, nil
}
