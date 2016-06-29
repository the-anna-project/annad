package clg

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

var methodNotFoundError = errgo.New("method not found")

// IsMethodNotFound asserts methodNotFoundError.
func IsMethodNotFound(err error) bool {
	return errgo.Cause(err) == methodNotFoundError
}

var invalidCLGExecutionError = errgo.New("invalid CLG execution")

// IsInvalidCLGExecution asserts invalidCLGExecutionError.
func IsInvalidCLGExecution(err error) bool {
	return errgo.Cause(err) == invalidCLGExecutionError
}
