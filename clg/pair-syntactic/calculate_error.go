package pairsyntactic

import (
	"github.com/juju/errgo"
)

var invalidFeatureKeyError = errgo.New("invalid feature key")

// IsInvalidFeatureKey asserts invalidFeatureKeyError.
func IsInvalidFeatureKey(err error) bool {
	return errgo.Cause(err) == invalidFeatureKeyError
}
