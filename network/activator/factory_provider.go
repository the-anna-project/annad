package activator

import (
	"github.com/xh3b4sd/anna/spec"
)

func (a *activator) Factory() spec.FactoryCollection {
	return a.FactoryCollection
}
