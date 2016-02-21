package impulse

import (
	"github.com/xh3b4sd/anna/spec"
)

func (i *impulse) MarshalJSON() ([]byte, error) {
	i.Log.WithTags(spec.Tags{L: "D", O: i, T: nil, V: 15}, "call MarshalJSON")
	return nil, nil
}
