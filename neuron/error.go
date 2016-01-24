package neuron

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var invalidNetworkRelationError = errgo.New("invalid network relation")

// IsInvalidNetworkRelation checks for the given error to be
// invalidNetworkRelationError. This error describes the issue of too many
// networks associated with a network. By design a network can only be related
// to one single network. If there are multiple networks detected, this error
// is returned.
func IsInvalidNetworkRelation(err error) bool {
	return errgo.Cause(err) == invalidNetworkRelationError
}

var neuronNotFoundError = errgo.New("neuron not found")

func IsNeuronNotFound(err error) bool {
	return errgo.Cause(err) == neuronNotFoundError
}

var invalidImpulseStateError = errgo.New("invalid impulse state")

// IsInvalidImpulseState checks for the given error to be
// invalidImpulseStateError. This error describes the issue of an unknown state
// of an impulse. By design an impulse can only have one of the following
// states:
//
//   "":             The impulse is going to be processed ASAP.
//   "in-progress":  Actions and changes are still applied to the impulse.
//   "finished":     No more action or change is applied to the impulse.
//
//  If there is a unknown state detected, this error is returned.
func IsInvalidImpulseState(err error) bool {
	return errgo.Cause(err) == invalidImpulseStateError
}
