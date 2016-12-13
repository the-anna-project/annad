package context

import (
	"encoding/json"
)

func (c *context) MarshalJSON() ([]byte, error) {
	type ContextClone context

	b, err := json.Marshal(&struct {
		*ContextClone
	}{
		ContextClone: (*ContextClone)(c),
	})
	if err != nil {
		return nil, maskAny(err)
	}

	return b, nil
}

func (c *context) UnmarshalJSON(b []byte) error {
	type ContextClone context

	aux := &struct {
		*ContextClone
	}{
		ContextClone: (*ContextClone)(c),
	}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
