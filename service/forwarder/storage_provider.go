package forwarder

import (
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

func (s *service) Storage() storagespec.Collection {
	return s.StorageCollection
}
