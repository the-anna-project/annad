package memory

import (
	"github.com/juju/errgo"

	"github.com/xh3b4sd/anna/storage/redis"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

// IsNotFound redirects to redis.IsNotFound.
var IsNotFound = redis.IsNotFound
