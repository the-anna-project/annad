package storage

import (
	"github.com/garyburd/redigo/redis"
)

type RedisDialConfig struct {
	// Addr representes the address used to connect to a redis server.
	Addr string
}

func DefaultRedisDialConfig() RedisDialConfig {
	newConfig := RedisDialConfig{
		Addr: "127.0.0.1:6379",
	}

	return newConfig
}

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
