package storage

import (
	"github.com/xh3b4sd/anna/storage/memory"
	"github.com/xh3b4sd/anna/storage/redis"
)

func IsNotFound(err error) bool {
	return redis.IsNotFound(err) || memory.IsNotFound(err)
}
