package strategy

import (
	"encoding/json"
)

// strategyClone is for making use of the stdlib json implementation. The
// strategy object implements its own marshaler and unmarshaler but only to
// provide json implementations for spec.Strategy. Note, not redirecting the
// type will cause infinite recursion.
type strategyClone strategy

func (s *strategy) MarshalJSON() ([]byte, error) {
	newStrategy := strategyClone(*s)

	raw, err := json.Marshal(newStrategy)
	if err != nil {
		return nil, maskAny(err)
	}

	return raw, nil
}

func (s *strategy) UnmarshalJSON(b []byte) error {
	newStrategy := strategyClone{}

	err := json.Unmarshal(b, &newStrategy)
	if err != nil {
		return maskAny(err)
	}

	s.CLGNames = newStrategy.CLGNames
	s.ID = newStrategy.ID
	s.Requestor = newStrategy.Requestor
	s.Type = newStrategy.Type

	return nil
}
