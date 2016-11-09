package activator

import (
	"encoding/json"

	objectspec "github.com/xh3b4sd/anna/object/spec"
)

func (s *service) persistQueue(key string, queue []objectspec.NetworkPayload) error {
	b, err := json.Marshal(queue)
	if err != nil {
		return maskAny(err)
	}
	err = s.Storage().General().Set(key, string(b))
	if err != nil {
		return maskAny(err)
	}

	return nil
}
