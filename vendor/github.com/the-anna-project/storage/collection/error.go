package collection

import "github.com/the-anna-project/storage/service/redis"

// IsNotFound combines IsNotFound error matchers of all storage
// implementations. IsNotFound should thus be used for error handling wherever
// spec.Storage is dealt with.
func IsNotFound(err error) bool {
	return redis.IsNotFound(err)
}
