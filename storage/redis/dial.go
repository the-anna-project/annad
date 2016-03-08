package redisstorage

import (
	"github.com/garyburd/redigo/redis"
	"github.com/rafaeljusto/redigomock"
)

// redis

// RedisDialConfig represents the configuration used to create a new redis
// dialer.
type RedisDialConfig struct {
	// Addr representes the address used to connect to a redis server.
	Addr string
}

// DefaultRedisDialConfig provides a default configuration to create a new
// redis dialer by best effort.
func DefaultRedisDialConfig() RedisDialConfig {
	newConfig := RedisDialConfig{
		Addr: "127.0.0.1:6379",
	}

	return newConfig
}

// NewRedisDial creates a new configured redis dialer.
func NewRedisDial(config RedisDialConfig) func() (redis.Conn, error) {
	newDial := func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", config.Addr)
		if err != nil {
			return nil, err
		}

		return c, err
	}

	return newDial
}

// mock

type mockDialConfig struct {
	RedisConn redis.Conn
}

func defaultMockDialConfig() mockDialConfig {
	newConfig := mockDialConfig{
		RedisConn: redigomock.NewConn(),
	}

	return newConfig
}

func newMockDial(config mockDialConfig) func() (redis.Conn, error) {
	newDial := func() (redis.Conn, error) {
		return config.RedisConn, nil
	}

	return newDial
}
