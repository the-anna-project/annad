package profile

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

var invalidCLGError = errgo.New("invalid clg")

// IsInvalidCLG asserts invalidCLGError.
func IsInvalidCLG(err error) bool {
	return errgo.Cause(err) == invalidCLGError
}

var clgProfileNotFoundError = errgo.New("clg profile not found")

// IsCLGProfileNotFound asserts clgProfileNotFoundError.
func IsCLGProfileNotFound(err error) bool {
	return errgo.Cause(err) == clgProfileNotFoundError
}

var clgBodyNotFoundError = errgo.New("clg body not found")

// IsCLGBodyNotFound asserts clgBodyNotFoundError.
func IsCLGBodyNotFound(err error) bool {
	return errgo.Cause(err) == clgBodyNotFoundError
}
