package output

import (
	"github.com/juju/errgo"
)

var expectationNotMetError = errgo.New("expectation not met")

// IsExpectationNotMet asserts expectationNotMetError.
func IsExpectationNotMet(err error) bool {
	return errgo.Cause(err) == expectationNotMetError
}
