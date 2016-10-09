package forwarder

import (
	"github.com/xh3b4sd/anna/spec"
)

func (f *forwarder) Factory() spec.FactoryCollection {
	return f.FactoryCollection
}
