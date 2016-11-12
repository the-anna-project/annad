package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

// PoolConfig represents the configuration used to create a new redis
// pool.
type PoolConfig struct {
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

// DefaultPoolConfig provides a default configuration to create a new
// redis pool by best effort.
func DefaultPoolConfig() PoolConfig {
	newConfig := PoolConfig{
		MaxIdle:     100,
		MaxActive:   100,
		IdleTimeout: 5 * time.Minute,
		Dial:        NewDial(DefaultDialConfig()),
	}

	return newConfig
}

// NewPool creates a new configured redis pool.
func NewPool(config PoolConfig) *redis.Pool {
	newPool := &redis.Pool{
		MaxIdle:     config.MaxIdle,
		MaxActive:   config.MaxActive,
		IdleTimeout: config.IdleTimeout,
		Dial:        config.Dial,
	}

	return newPool
}
