package corenet

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var networkNotFoundError = errgo.New("network not found")

func IsNetworkNotFound(err error) bool {
	return errgo.Cause(err) == networkNotFoundError
}
