package forwarder

import (
	"github.com/xh3b4sd/anna/spec"
)

func (f *forwarder) Service() spec.ServiceCollection {
	return f.ServiceCollection
}
