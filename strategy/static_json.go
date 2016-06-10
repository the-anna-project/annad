package strategy

import (
	"encoding/json"
)

// staticClone is for making use of the stdlib json implementation. The static
// object implements its own marshaler and unmarshaler but only to provide json
// implementations for spec.Strategy. Note, not redirecting the type will cause
// infinite recursion.
type staticClone static

func (s *static) MarshalJSON() ([]byte, error) {
	newStrategy := staticClone(*s)

	raw, err := json.Marshal(newStrategy)
	if err != nil {
		return nil, maskAny(err)
	}

	return raw, nil
}

func (s *static) UnmarshalJSON(b []byte) error {
	newStrategy := staticClone{}

	err := json.Unmarshal(b, &newStrategy)
	if err != nil {
		return maskAny(err)
	}

	s.ID = newStrategy.ID
	s.Argument = newStrategy.Argument
	s.Type = newStrategy.Type

	return nil
}
