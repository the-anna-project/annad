package textinterface

import (
	"github.com/xh3b4sd/anna/spec"
)

func (ti *textInterface) UnmarshalJSON(bytes []byte) error {
	ti.Log.WithTags(spec.Tags{L: "D", O: ti, T: nil, V: 15}, "call UnmarshalJSON")
	return nil
}
