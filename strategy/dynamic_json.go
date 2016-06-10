package strategy

import (
	"encoding/json"
)

// dynamicClone is for making use of the stdlib json implementation. The dynamic
// object implements its own marshaler and unmarshaler but only to provide json
// implementations for spec.Strategy. Note, not redirecting the type will cause
// infinite recursion.
type dynamicClone dynamic

func (d *dynamic) MarshalJSON() ([]byte, error) {
	newStrategy := dynamicClone(*d)

	raw, err := json.Marshal(newStrategy)
	if err != nil {
		return nil, maskAny(err)
	}

	return raw, nil
}

func (d *dynamic) UnmarshalJSON(b []byte) error {
	newStrategy := dynamicClone{}

	err := json.Unmarshal(b, &newStrategy)
	if err != nil {
		return maskAny(err)
	}

	d.ID = newStrategy.ID
	d.Nodes = newStrategy.Nodes
	d.Root = newStrategy.Root
	d.Type = newStrategy.Type

	return nil
}
