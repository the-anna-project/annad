package gateway

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var gatewayClosedError = errgo.New("gateway closed")

// IsGatewayClosed asserts gatewayClosedError.
func IsGatewayClosed(err error) bool {
	return errgo.Cause(err) == gatewayClosedError
}

var signalCanceledError = errgo.New("signal canceled")

// IsSignalCanceled asserts signalCanceledError.
func IsSignalCanceled(err error) bool {
	return errgo.Cause(err) == signalCanceledError
}
