package common

import (
	"github.com/xh3b4sd/anna/spec"
)

func MustObjectToCore(object spec.Object) spec.Core {
	if i, ok := object.(spec.Core); ok {
		return i
	}

	panic(objectNotCoreError)
}
