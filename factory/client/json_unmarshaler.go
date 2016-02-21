package factoryclient

import (
	"github.com/xh3b4sd/anna/spec"
)

func (fc *factoryClient) UnmarshalJSON(bytes []byte) error {
	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call UnmarshalJSON")
	return nil
}
