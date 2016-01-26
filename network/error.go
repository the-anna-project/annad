package network

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var stateNotFoundError = errgo.New("state not found")

func IsStateNotFound(err error) bool {
	return errgo.Cause(err) == stateNotFoundError
}
