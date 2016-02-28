package impulse

import (
	"github.com/xh3b4sd/anna/spec"
)

func (i *impulse) UnmarshalJSON(bytes []byte) error {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 15}, "call UnmarshalJSON")
	return nil
}
