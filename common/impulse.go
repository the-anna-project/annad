package common

import (
	"github.com/xh3b4sd/anna/spec"
)

const (
	ImpulseIDKey = "impulse-id"
)

func MustObjectToImpulse(object spec.Object) spec.Impulse {
	if i, ok := object.(spec.Impulse); ok {
		return i
	}

	panic(objectNotImpulseError)
}

func GetInitImpulseCopy(key string, object spec.Object) (spec.Impulse, error) {
	objectState, err := object.GetState(InitStateKey)
	if err != nil {
		return nil, maskAny(err)
	}
	bytes, err := objectState.GetBytes(key)
	if err != nil {
		return nil, maskAny(err)
	}
	initImpulse, err := objectState.GetImpulseByID(spec.ObjectID(bytes))
	if err != nil {
		return nil, maskAny(err)
	}

	return initImpulse.Copy(), nil
}
