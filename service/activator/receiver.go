package activator

import (
	"encoding/json"

	objectspec "github.com/the-anna-project/spec/object"
)

func (s *service) persistQueue(key string, queue []objectspec.NetworkPayload) error {
	b, err := json.Marshal(queue)
	if err != nil {
		return maskAny(err)
	}
	err = s.Service().Storage().General().Set(key, string(b))
	if err != nil {
		return maskAny(err)
	}

	return nil
}
