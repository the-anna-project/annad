package argument

import (
	"fmt"

	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

func maskAnyf(err error, f string, v ...interface{}) error {
	if err == nil {
		return nil
	}

	f = fmt.Sprintf("%s: %s", err.Error(), f)
	newErr := errgo.WithCausef(nil, errgo.Cause(err), f, v...)
	newErr.(*errgo.Err).SetLocation(1)

	return newErr
}

var invalidConfigError = errgo.New("invalid config")

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return errgo.Cause(err) == invalidConfigError
}

var indexOutOfRangeError = errgo.New("index out of range")

// IsIndexOutOfRange asserts indexOutOfRangeError.
func IsIndexOutOfRange(err error) bool {
	return errgo.Cause(err) == indexOutOfRangeError
}

var invalidTypeError = errgo.New("invalid type")

// IsInvalidType asserts invalidTypeError.
func IsInvalidType(err error) bool {
	return errgo.Cause(err) == invalidTypeError
}
