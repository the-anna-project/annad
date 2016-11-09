package network

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

var clgNotFoundError = errgo.New("clg not found")

// IsCLGNotFound asserts clgNotFoundError.
func IsCLGNotFound(err error) bool {
	return errgo.Cause(err) == clgNotFoundError
}

var invalidInterfaceError = errgo.New("invalid interface")

// IsInvalidInterface asserts invalidInterfaceError.
func IsInvalidInterface(err error) bool {
	return errgo.Cause(err) == invalidInterfaceError
}

var invalidBehaviourIDError = errgo.New("invalid behaviour ID")

// IsInvalidBehaviourID asserts invalidBehaviourIDError.
func IsInvalidBehaviourID(err error) bool {
	return errgo.Cause(err) == invalidBehaviourIDError
}

var invalidCLGNameError = errgo.New("invalid CLG name")

// IsInvalidCLGName asserts invalidCLGNameError.
func IsInvalidCLGName(err error) bool {
	return errgo.Cause(err) == invalidCLGNameError
}

var invalidCLGTreeIDError = errgo.New("invalid CLG tree ID")

// IsInvalidCLGTreeID asserts invalidCLGTreeIDError.
func IsInvalidCLGTreeID(err error) bool {
	return errgo.Cause(err) == invalidCLGTreeIDError
}

var invalidInformationIDError = errgo.New("invalid information ID")

// IsInvalidInformationID asserts invalidInformationIDError.
func IsInvalidInformationID(err error) bool {
	return errgo.Cause(err) == invalidInformationIDError
}

var invalidNetworkPayloadError = errgo.New("invalid network payload")

// IsInvalidNetworkPayload asserts invalidNetworkPayloadError.
func IsInvalidNetworkPayload(err error) bool {
	return errgo.Cause(err) == invalidNetworkPayloadError
}

var workerCanceledError = errgo.New("worker canceled")

// IsWorkerCanceled asserts workerCanceledError.
func IsWorkerCanceled(err error) bool {
	return errgo.Cause(err) == workerCanceledError
}
