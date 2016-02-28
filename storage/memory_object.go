package storage

import (
	"github.com/xh3b4sd/anna/spec"
)

func (ms *memoryStorage) GetID() spec.ObjectID {
	ms.Mutex.Lock()
	defer ms.Mutex.Unlock()
	return ms.ID
}

func (ms *memoryStorage) GetType() spec.ObjectType {
	ms.Mutex.Lock()
	defer ms.Mutex.Unlock()
	return ms.Type
}
