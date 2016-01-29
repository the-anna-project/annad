package spec

import (
	"encoding/json"
)

type Network interface {
	json.Unmarshaler

	Object

	Trigger(imp Impulse) (Impulse, error)
}
