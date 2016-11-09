package readinformationid

import (
	"github.com/juju/errgo"
)

var invalidInformationIDError = errgo.New("invalid information ID")

// IsInvalidInformationID asserts invalidInformationIDError.
func IsInvalidInformationID(err error) bool {
	return errgo.Cause(err) == invalidInformationIDError
}
