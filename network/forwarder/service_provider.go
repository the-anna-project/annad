package forwarder

import (
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

func (f *forwarder) Service() servicespec.Collection {
	return f.ServiceCollection
}
