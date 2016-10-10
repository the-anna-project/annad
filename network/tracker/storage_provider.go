package tracker

import (
	"github.com/xh3b4sd/anna/spec"
)

func (t *tracker) Storage() spec.StorageCollection {
	return t.StorageCollection
}
