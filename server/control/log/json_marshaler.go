package logcontrol

import (
	"github.com/xh3b4sd/anna/spec"
)

func (lc *logControl) MarshalJSON() ([]byte, error) {
	lc.Log.WithTags(spec.Tags{L: "D", O: lc, T: nil, V: 15}, "call MarshalJSON")
	return nil, nil
}
