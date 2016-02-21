package factoryclient

import (
	"github.com/xh3b4sd/anna/spec"
)

func (fc *factoryClient) MarshalJSON() ([]byte, error) {
	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call MarshalJSON")
	return nil, nil
}
