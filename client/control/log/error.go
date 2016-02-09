package logcontrol

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

var invalidAPIResponseError = errgo.New("invalid api response")

// IsInvalidAPIResponse checks for the given error to be
// invalidAPIResponseError. This error is returned in case there is an
// unexpected API response received from the server.
func IsInvalidAPIResponse(err error) bool {
	return errgo.Cause(err) == invalidAPIResponseError
}
