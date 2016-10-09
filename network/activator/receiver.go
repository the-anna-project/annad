package activator

import (
	"encoding/json"

	"github.com/xh3b4sd/anna/spec"
)

func (a *activator) persistQueue(key string, queue []spec.NetworkPayload) error {
	b, err := json.Marshal(queue)
	if err != nil {
		return maskAny(err)
	}
	err = a.Storage().General().Set(key, string(b))
	if err != nil {
		return maskAny(err)
	}

	return nil
}
