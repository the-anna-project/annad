package forwarder

import (
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

func (f *forwarder) Storage() storagespec.Collection {
	return f.StorageCollection
}
