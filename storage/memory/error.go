package memorystorage

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var keyNotFoundError = errgo.New("key not found")

func IsKeyNotFound(err error) bool {
	return errgo.Cause(err) == keyNotFoundError
}
