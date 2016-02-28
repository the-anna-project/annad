package log

import (
	"github.com/xh3b4sd/anna/spec"
)

func (l *log) MarshalJSON() ([]byte, error) {
	l.WithTags(spec.Tags{L: "D", O: l, T: nil, V: 15}, "call MarshalJSON")
	return nil, nil
}
