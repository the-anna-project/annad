package memory

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var notFoundError = errgo.New("not found")

// IsNotFound asserts notFoundError.
func IsNotFound(err error) bool {
	return errgo.Cause(err) == notFoundError
}
