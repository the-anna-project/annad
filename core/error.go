package core

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var connectionNotFoundError = errgo.New("connection not found")

func IsConnectionNotFound(err error) bool {
	return errgo.Cause(err) == connectionNotFoundError
}
