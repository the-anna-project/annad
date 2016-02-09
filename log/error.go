package log

import (
	"fmt"

	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

func maskAnyf(err error, f string, v ...interface{}) error {
	f = fmt.Sprintf("%s: %s", err.Error(), f)
	newErr := errgo.WithCausef(err, errgo.Cause(err), f, v...)

	if e, _ := newErr.(*errgo.Err); e != nil {
		e.SetLocation(1)
		return e
	}

	return err
}

var invalidLogLevelError = errgo.New("invalid log level")

func IsInvalidLogLevel(err error) bool {
	return errgo.Cause(err) == invalidLogLevelError
}

var invalidLogObjectError = errgo.New("invalid log object")

func IsInvalidLogObject(err error) bool {
	return errgo.Cause(err) == invalidLogObjectError
}
