package spec

import (
	"reflect"
)

// TODO implement

// Output represents the result of some calculations.
type Output interface {
	// GetType returns the outputs reflect.Type.
	GetType() reflect.Type

	// Match checks whether the current output matches the given expectation.
	Match(expectation Expectation) (bool, error)

	// String returns the outputs string representation.
	String() string
}
