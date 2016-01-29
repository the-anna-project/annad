package spec

import (
	"encoding/json"
)

type Impulse interface {
	json.Unmarshaler

	Object

	WalkThrough(neu Neuron) (Impulse, Neuron, error)
}
