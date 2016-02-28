package storage

import (
	"github.com/garyburd/redigo/redis"
	"github.com/rafaeljusto/redigomock"
)

type MemoryDialConfig struct {
	RedisConn redis.Conn
}

func DefaultMemoryDialConfig() MemoryDialConfig {
	newConfig := MemoryDialConfig{
		RedisConn: redigomock.NewConn(),
	}

	return newConfig
}

func NewMemoryDial(config MemoryDialConfig) func() (redis.Conn, error) {
	newDial := func() (redis.Conn, error) {
		return config.RedisConn, nil
	}

	return newDial
}
