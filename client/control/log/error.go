package log

import (
	"encoding/json"
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

var invalidAPIResponseError = errgo.New("invalid api response")

// IsInvalidAPIResponse asserts invalidAPIResponseError.
func IsInvalidAPIResponse(err error) bool {
	return errgo.Cause(err) == invalidAPIResponseError
}

// IsUnsupportedType asserts json.UnsupportedTypeError.
func IsUnsupportedType(err error) bool {
	if _, ok := errgo.Cause(err).(*json.UnsupportedTypeError); ok {
		return true
	}

	return false
}
