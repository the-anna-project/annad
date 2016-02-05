package factoryserver

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

var invalidFactoryGatewayRequestError = errgo.New("invalid factory gateway request")

func IsInvalidFactoryGatewayRequest(err error) bool {
	return errgo.Cause(err) == invalidFactoryGatewayRequestError
}
