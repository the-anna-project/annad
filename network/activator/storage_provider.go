package activator

import (
	"github.com/xh3b4sd/anna/spec"
)

func (a *activator) Storage() spec.StorageCollection {
	return a.StorageCollection
}
