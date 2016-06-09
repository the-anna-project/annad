package spec

import (
	"encoding/json"
	"reflect"
)

// Strategy represents an executable and permutable CLG tree structure.
type Strategy interface {
	Execute() ([]reflect.Value, error)

	GetRoot() CLG

	GetNodes() []Strategy

	json.Marshaler

	Object

	json.Unmarshaler

	SetNode(indizes []int, node Strategy) error

	Validate() error
}
