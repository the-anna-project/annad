package client

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var invalidAPIResponseError = errgo.New("invalid api response")

// IsInvalidAPIResponse checks for the given error to be
// invalidAPIResponseError. This error is returned in case there is an
// unexpected API response received from the server.
func IsInvalidAPIResponse(err error) bool {
	return errgo.Cause(err) == invalidAPIResponseError
}
