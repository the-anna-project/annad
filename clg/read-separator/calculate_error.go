package readseparator

import (
	"github.com/juju/errgo"
)

var invalidBehaviorIDError = errgo.New("invalid behavior ID")

// IsInvalidBehaviorID asserts invalidBehaviorIDError.
func IsInvalidBehaviorID(err error) bool {
	return errgo.Cause(err) == invalidBehaviorIDError
}

var invalidFeatureKeyError = errgo.New("invalid feature key")

// IsInvalidFeatureKey asserts invalidFeatureKeyError.
func IsInvalidFeatureKey(err error) bool {
	return errgo.Cause(err) == invalidFeatureKeyError
}
