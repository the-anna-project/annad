package connection

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

var invalidPeerError = errgo.New("invalid config")

// IsInvalidPeer asserts invalidPeerError.
func IsInvalidPeer(err error) bool {
	return errgo.Cause(err) == invalidPeerError
}

var connectionNotFoundError = errgo.New("connection not found")

// IsConnectionNotFound asserts connectionNotFoundError.
func IsConnectionNotFound(err error) bool {
	return errgo.Cause(err) == connectionNotFoundError
}
