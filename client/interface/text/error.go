package textinterface

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

func maskAnyWithCause(underlying, cause error) error {
	err := maskAny(errgo.WithCausef(underlying, cause, ""))

	if e, _ := err.(*errgo.Err); e != nil {
		e.SetLocation(1)
	}

	return err
}

var invalidAPIResponseError = errgo.New("invalid api response")

// IsInvalidAPIResponse checks for the given error to be
// invalidAPIResponseError. This error is returned in case there is an
// unexpected API response received from the server.
func IsInvalidAPIResponse(err error) bool {
	return errgo.Cause(err) == invalidAPIResponseError
}
