package memoryfilesystem

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var noSuchFileOrDirectoryError = errgo.New("no such file or directory")

// IsNoSuchFileOrDirectoryError asserts noSuchFileOrDirectoryError.
func IsNoSuchFileOrDirectoryError(err error) bool {
	return errgo.Cause(err) == noSuchFileOrDirectoryError
}
