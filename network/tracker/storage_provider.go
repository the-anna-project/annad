package tracker

import (
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

func (t *tracker) Storage() storagespec.Collection {
	return t.StorageCollection
}
