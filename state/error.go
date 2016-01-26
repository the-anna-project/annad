package state

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var invalidCoreRelationError = errgo.New("invalid core relation")

// IsInvalidCoreRelation checks for the given error to be
// invalidCoreRelationError. This error is returned in case there is not
// exactly one core. By design there can only be one core. If there is none, or
// too many, this error is returned.
func IsInvalidCoreRelation(err error) bool {
	return errgo.Cause(err) == invalidCoreRelationError
}

var networkNotFoundError = errgo.New("network not found")

// IsNetworkNotFound checks for the given error to be networkNotFoundError.
// This error is returned in case there is no network as required.
func IsNetworkNotFound(err error) bool {
	return errgo.Cause(err) == networkNotFoundError
}

var neuronNotFoundError = errgo.New("neuron not found")

// IsNeuronNotFound checks for the given error to be neuronNotFoundError.
// This error is returned in case there is no neuron as required.
func IsNeuronNotFound(err error) bool {
	return errgo.Cause(err) == neuronNotFoundError
}

var bytesNotFoundError = errgo.New("bytes not found")

// IsBytesNotFound checks for the given error to be bytesNotFoundError.
// This error is returned in case there is no bytes as required.
func IsBytesNotFound(err error) bool {
	return errgo.Cause(err) == bytesNotFoundError
}

var coreNotFoundError = errgo.New("core not found")

// IsCoreNotFound checks for the given error to be coreNotFoundError.
// This error is returned in case there is no core as required.
func IsCoreNotFound(err error) bool {
	return errgo.Cause(err) == coreNotFoundError
}

var impulseNotFoundError = errgo.New("impulse not found")

// IsImpulseNotFound checks for the given error to be impulseNotFoundError.
// This error is returned in case there is no impulse as required.
func IsImpulseNotFound(err error) bool {
	return errgo.Cause(err) == impulseNotFoundError
}
