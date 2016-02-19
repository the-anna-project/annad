package textinterface

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

var invalidAPIResponseError = errgo.New("invalid api response")

// IsInvalidAPIResponse checks for the given error to be
// invalidAPIResponseError. This error is returned in case there is an
// unexpected API response received from the server.
func IsInvalidAPIResponse(err error) bool {
	return errgo.Cause(err) == invalidAPIResponseError
}
