package spec

import (
	"reflect"
)

// TODO implement

// Input represents the input provided when requesting some calculations.
type Input interface {
	// GetType returns the outputs reflect.Type.
	GetType() reflect.Type

	// String returns the inputs string representation.
	String() string
}
