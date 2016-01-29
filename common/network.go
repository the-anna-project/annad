package common

import (
	"github.com/xh3b4sd/anna/spec"
)

func MustObjectToNetwork(object spec.Object) spec.Network {
	if i, ok := object.(spec.Network); ok {
		return i
	}

	panic(objectNotNetworkError)
}
