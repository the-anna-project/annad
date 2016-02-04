package common

import (
	"github.com/xh3b4sd/anna/spec"
)

func MustObjectToImpulse(object spec.Object) spec.Impulse {
	if i, ok := object.(spec.Impulse); ok {
		return i
	}

	panic(objectNotImpulseError)
}
