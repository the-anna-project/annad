package forwarder

import (
	"github.com/xh3b4sd/anna/spec"
)

func (f *forwarder) Storage() spec.StorageCollection {
	return f.StorageCollection
}
