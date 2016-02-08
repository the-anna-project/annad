package log

import (
	"fmt"

	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

func maskAnyf(cause error, f string, v ...interface{}) error {
	f = fmt.Sprintf("%s: %s", cause.Error(), f)
	err := maskAny(errgo.WithCausef(nil, cause, f, v...))

	if e, _ := err.(*errgo.Err); e != nil {
		e.SetLocation(1)
	}

	return err
}

var invalidLogLevelError = errgo.New("invalid log level")

func IsInvalidLogLevel(err error) bool {
	return errgo.Cause(err) == invalidLogLevelError
}
