package scheduler

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var jobNotFoundError = errgo.New("job not found")

// IsJobNotFound asserts jobNotFoundError.
func IsJobNotFound(err error) bool {
	return errgo.Cause(err) == jobNotFoundError
}
