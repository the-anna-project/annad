package tracker

import (
	"github.com/xh3b4sd/anna/spec"
)

func (t *tracker) Service() spec.ServiceCollection {
	return t.ServiceCollection
}
