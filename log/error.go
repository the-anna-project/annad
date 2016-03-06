package log

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

var invalidLogLevelError = errgo.New("invalid log level")

// IsInvalidLogLevel asserts invalidLogLevelError.
func IsInvalidLogLevel(err error) bool {
	return errgo.Cause(err) == invalidLogLevelError
}

var invalidLogObjectError = errgo.New("invalid log object")

// IsInvalidLogObject asserts invalidLogObjectError.
func IsInvalidLogObject(err error) bool {
	return errgo.Cause(err) == invalidLogObjectError
}
