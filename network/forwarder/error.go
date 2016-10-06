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

var networkPayloadsNotFoundError = errgo.New("network payloads not found")

// IsNetworkPayloadsNotFound asserts networkPayloadsNotFoundError.
func IsNetworkPayloadsNotFound(err error) bool {
	return errgo.Cause(err) == networkPayloadsNotFoundError
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
