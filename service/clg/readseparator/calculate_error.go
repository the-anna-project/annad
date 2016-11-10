package readseparator

import (
	"github.com/juju/errgo"
)

var invalidBehaviourIDError = errgo.New("invalid behaviour ID")

// IsInvalidBehaviourID asserts invalidBehaviourIDError.
func IsInvalidBehaviourID(err error) bool {
	return errgo.Cause(err) == invalidBehaviourIDError
}

var invalidFeatureKeyError = errgo.New("invalid feature key")

// IsInvalidFeatureKey asserts invalidFeatureKeyError.
func IsInvalidFeatureKey(err error) bool {
	return errgo.Cause(err) == invalidFeatureKeyError
}
