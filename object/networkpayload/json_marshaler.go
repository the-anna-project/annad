package networkpayload

import (
	"encoding/json"
)

func (np *networkPayload) MarshalJSON() ([]byte, error) {
	type NetworkPayloadClone networkPayload

	b, err := json.Marshal(&struct {
		*NetworkPayloadClone
	}{
		NetworkPayloadClone: (*NetworkPayloadClone)(np),
	})
	if err != nil {
		return nil, maskAny(err)
	}

	return b, nil
}
