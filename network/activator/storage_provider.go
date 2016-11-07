package activator

import (
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

func (a *activator) Storage() storagespec.Collection {
	return a.StorageCollection
}
