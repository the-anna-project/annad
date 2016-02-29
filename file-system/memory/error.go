package memoryfilesystem

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var noSuchFileOrDirectoryError = errgo.New("no such file or directory")

// IsNoSuchFileOrDirectoryError checks for the given error to be
// noSuchFileOrDirectoryError. This error is returned in case there cannot
// any file be found as requested.
func IsNoSuchFileOrDirectoryError(err error) bool {
	return errgo.Cause(err) == noSuchFileOrDirectoryError
}
