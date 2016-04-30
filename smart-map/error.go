package smartmap

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

var unsupportedTypeError = errgo.New("unsupported type")

// IsUnsupportedType asserts unsupportedTypeError.
func IsUnsupportedType(err error) bool {
	return errgo.Cause(err) == unsupportedTypeError
}

var keyNotFoundError = errgo.New("key not found")

// IsKeyNotFound asserts keyNotFoundError.
func IsKeyNotFound(err error) bool {
	return errgo.Cause(err) == keyNotFoundError
}

var wrongTypeError = errgo.New("wrong type")

// IsWrongType asserts wrongTypeError.
func IsWrongType(err error) bool {
	return errgo.Cause(err) == wrongTypeError
}

var prefixNotFoundError = errgo.New("wrong type")

// IsPrefixNotFound asserts prefixNotFoundError.
func IsPrefixNotFound(err error) bool {
	return errgo.Cause(err) == prefixNotFoundError
}
