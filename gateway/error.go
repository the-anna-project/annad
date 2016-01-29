package gateway

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var gatewayClosedError = errgo.New("gateway closed")

func IsGatewayClosed(err error) bool {
	return errgo.Cause(err) == gatewayClosedError
}

var signalCanceledError = errgo.New("signal canceled")

func IsSignalCanceled(err error) bool {
	return errgo.Cause(err) == signalCanceledError
}

var bytesNotFoundError = errgo.New("bytes not found")

// IsBytesNotFound checks for the given error to be bytesNotFoundError.
// This error is returned in case there is no bytes as required.
func IsBytesNotFound(err error) bool {
	return errgo.Cause(err) == bytesNotFoundError
}

var objectNotFoundError = errgo.New("object not found")

// IsObjectNotFound checks for the given error to be objectNotFoundError.
// This error is returned in case there is no object as required.
func IsObjectNotFound(err error) bool {
	return errgo.Cause(err) == objectNotFoundError
}
