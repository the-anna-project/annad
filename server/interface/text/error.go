package textinterface

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

var invalidRequestError = errgo.New("invalid request")

func IsInvalidRequest(err error) bool {
	return errgo.Cause(err) == invalidRequestError
}
