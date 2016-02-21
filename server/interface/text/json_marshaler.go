package textinterface

import (
	"github.com/xh3b4sd/anna/spec"
)

func (ti *textInterface) MarshalJSON() ([]byte, error) {
	ti.Log.WithTags(spec.Tags{L: "D", O: ti, T: nil, V: 15}, "call MarshalJSON")
	return nil, nil
}
