package output

import (
	"github.com/juju/errgo"
)

var expectationNotMetError = errgo.New("expectation not met")

// IsExpectationNotMet asserts expectationNotMetError.
func IsExpectationNotMet(err error) bool {
	return errgo.Cause(err) == expectationNotMetError
}

var invalidBehaviourIDError = errgo.New("invalid behaviour ID")

// IsInvalidBehaviourID asserts invalidBehaviourIDError.
func IsInvalidBehaviourID(err error) bool {
	return errgo.Cause(err) == invalidBehaviourIDError
}

var invalidInformationIDError = errgo.New("invalid information ID")

// IsInvalidInformationID asserts invalidInformationIDError.
func IsInvalidInformationID(err error) bool {
	return errgo.Cause(err) == invalidInformationIDError
}

var invalidCLGTreeIDError = errgo.New("invalid CLG tree ID")

// IsInvalidCLGTreeID asserts invalidCLGTreeIDError.
func IsInvalidCLGTreeID(err error) bool {
	return errgo.Cause(err) == invalidCLGTreeIDError
}
