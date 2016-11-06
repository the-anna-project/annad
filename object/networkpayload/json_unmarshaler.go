package networkpayload

import (
	"encoding/json"
)

func (np *networkPayload) UnmarshalJSON(b []byte) error {
	type NetworkPayloadClone networkPayload

	aux := &struct {
		*NetworkPayloadClone
	}{
		NetworkPayloadClone: (*NetworkPayloadClone)(np),
	}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
