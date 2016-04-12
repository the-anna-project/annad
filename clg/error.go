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

var tooManyArgumentsError = errgo.New("not enough arguments")

// IsTooManyArguments asserts tooManyArgumentsError.
func IsTooManyArguments(err error) bool {
	return errgo.Cause(err) == tooManyArgumentsError
}

var wrongArgumentTypeError = errgo.New("wrong argument type")

// IsWrongArgumentType asserts wrongArgumentTypeError.
func IsWrongArgumentType(err error) bool {
	return errgo.Cause(err) == wrongArgumentTypeError
}

var methodNotFoundError = errgo.New("method not found")

// IsMethodNotFound asserts methodNotFoundError.
func IsMethodNotFound(err error) bool {
	return errgo.Cause(err) == methodNotFoundError
}

var indexOutOfRangeError = errgo.New("index out of range")

// IsIndexOutOfRange asserts indexOutOfRangeError.
func IsIndexOutOfRange(err error) bool {
	return errgo.Cause(err) == indexOutOfRangeError
}

var negativeIntError = errgo.New("negative integer")

// IsNegativeInt asserts negativeIntError.
func IsNegativeInt(err error) bool {
	return errgo.Cause(err) == negativeIntError
}

var invalidDividerError = errgo.New("invalid divider")

// IsInvalidDivider asserts invalidDividerError.
func IsInvalidDivider(err error) bool {
	return errgo.Cause(err) == invalidDividerError
}
