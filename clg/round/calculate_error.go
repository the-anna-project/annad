package round

import (
	"strconv"

	"github.com/juju/errgo"
)

// IsInvalidConfig asserts invalidConfigError.
func IsParseFloatSyntax(err error) bool {
	cause := errgo.Cause(err)

	if e, ok := cause.(*strconv.NumError); ok {
		if e.Err == strconv.ErrSyntax {
			return true
		}
	}

	return false
}
