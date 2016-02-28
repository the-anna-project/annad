package storage

import (
	"sync"

	"github.com/garyburd/redigo/redis"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

type RedisStorageConfig struct {
	Log  spec.Log
	Pool *redis.Pool
}

func DefaultRedisStorageConfig() RedisStorageConfig {
	newConfig := RedisStorageConfig{
		Log:  log.NewLog(log.DefaultConfig()),
		Pool: NewRedisPool(DefaultRedisPoolConfig()),
	}

	return newConfig
}

func NewRedisStorage(config RedisStorageConfig) spec.Storage {
	newStorage := &redisStorage{
		ID:                 id.NewObjectID(id.Hex128),
		Mutex:              sync.Mutex{},
		RedisStorageConfig: config,
		Type:               common.ObjectType.RedisStorage,
	}

	return newStorage
}

type redisStorage struct {
	ID    spec.ObjectID
	Mutex sync.Mutex `json:"-"`
	RedisStorageConfig
	Type spec.ObjectType
}

func (rs *redisStorage) Get(key string) (string, error) {
	conn := rs.Pool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", maskAny(err)
	}

	return value, nil
}

func (rs *redisStorage) GetElementsByScore(key string, score float32, maxElements int) ([]string, error) {
	conn := rs.Pool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("ZREVRANGEBYSCORE", key, score, score, "LIMIT", 0, maxElements))
	if err != nil {
		return nil, maskAny(err)
	}

	newList := []string{}
	for _, v := range values {
		newList = append(newList, v.(string))
	}

	return newList, nil
}

func (rs *redisStorage) GetHighestElementScore(key string) (string, float32, error) {
	conn := rs.Pool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("ZREVRANGE", key, 0, 0, "WITHSCORES"))
	if err != nil {
		return "", 0, maskAny(err)
	}

	if len(values) != 1 {
		return "", 0, maskAny(err)
	}

	return values[0].(string), values[1].(float32), nil
}

func (rs *redisStorage) Set(key, value string) error {
	conn := rs.Pool.Get()
	defer conn.Close()

	reply, err := redis.String(conn.Do("SET", key, value))
	if err != nil {
		return maskAny(err)
	}

	if reply != "OK" {
		return maskAnyf(queryExecutionFailedError, "SET not executed correctly")
	}

	return nil
}
