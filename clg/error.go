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

var notEnoughArgumentsError = errgo.New("not enough arguments")

// IsNotEnoughArguments asserts notEnoughArgumentsError.
func IsNotEnoughArguments(err error) bool {
	return errgo.Cause(err) == notEnoughArgumentsError
}

var wrongArgumentTypeError = errgo.New("wrong argument type")

// IsWrongArgumentType asserts wrongArgumentTypeError.
func IsWrongArgumentType(err error) bool {
	return errgo.Cause(err) == wrongArgumentTypeError
}
