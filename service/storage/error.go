package storage

import (
	"github.com/the-anna-project/annad/service/storage/redis"
)

// IsNotFound combines IsNotFound error matchers of all storage
// implementations. IsNotFound should thus be used for error handling wherever
// spec.Storage is dealt with.
func IsNotFound(err error) bool {
	return redis.IsNotFound(err)
}
