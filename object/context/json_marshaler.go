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
