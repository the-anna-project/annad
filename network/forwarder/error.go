package forwarder

import (
	"fmt"

	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

func maskAnyf(err error, f string, v ...interface{}) error {
	if err == nil {
		return nil
	}

	f = fmt.Sprintf("%s: %s", err.Error(), f)
	newErr := errgo.WithCausef(nil, errgo.Cause(err), f, v...)
	newErr.(*errgo.Err).SetLocation(1)

	return newErr
}

var invalidConfigError = errgo.New("invalid config")

// IsInvalidConfig asserts invalidConfigError.
func IsInvalidConfig(err error) bool {
	return errgo.Cause(err) == invalidConfigError
}

var behaviourIDsNotFoundError = errgo.New("behaviour IDs not found")

// IsBehaviourIDsNotFound asserts behaviourIDsNotFoundError.
func IsBehaviourIDsNotFound(err error) bool {
	return errgo.Cause(err) == behaviourIDsNotFoundError
}

var invalidBehaviorIDError = errgo.New("invalid behavior ID")

// IsInvalidBehaviorID asserts invalidBehaviorIDError.
func IsInvalidBehaviorID(err error) bool {
	return errgo.Cause(err) == invalidBehaviorIDError
}
