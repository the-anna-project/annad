package storage

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisPoolConfig struct {
	// MaxIdle is the allowed maximum number of idle connections in the pool.
	MaxIdle int

	// MaxActive is the allowed maximum number of connections allocated by the
	// pool at a given time.  When zero, there is no limit on the number of
	// connections in the pool.
	MaxActive int

	// Close connections after remaining idle for this duration. If the value
	// is zero, then idle connections are not closed. Applications should set
	// the timeout to a value less than the server's timeout.
	IdleTimeout time.Duration

	// Dial is an application supplied function for creating and configuring a
	// redis connection on demand.
	Dial func() (redis.Conn, error)
}

func DefaultRedisPoolConfig() RedisPoolConfig {
	newConfig := RedisPoolConfig{
		MaxIdle:     100,
		MaxActive:   100,
		IdleTimeout: 5 * time.Minute,
		Dial:        NewRedisDial(DefaultRedisDialConfig()),
	}

	return newConfig
}

func NewRedisPool(config RedisPoolConfig) *redis.Pool {
	newPool := &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: config.IdleTimeout,
		Dial:        config.Dial,
	}

	return newPool
}
