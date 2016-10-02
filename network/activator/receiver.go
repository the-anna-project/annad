package activator

import (
	"encoding/json"
)

func (a *activator) persistQueue(key string, queue []spec.NetworkPayload) error {
	raw, err = json.marshal(queue)
	if err != nil {
		return maskAny(err)
	}
	err := a.Storage().General().Set(key, string(raw))
	if err != nil {
		return maskAny(err)
	}

	return nil
}
