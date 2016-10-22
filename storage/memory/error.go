package memory

import (
	"github.com/juju/errgo"

	"github.com/xh3b4sd/anna/storage/redis"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var notFoundError = errgo.New("not found")

// sIsNotFound asserts notFoundError.
func IsNotFound(err error) bool {
	return errgo.Cause(err) == notFoundError || redis.IsNotFound(err)
}
