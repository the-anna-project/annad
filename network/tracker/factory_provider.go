package tracker

import (
	"github.com/xh3b4sd/anna/spec"
)

func (t *tracker) Factory() spec.FactoryCollection {
	return t.FactoryCollection
}
