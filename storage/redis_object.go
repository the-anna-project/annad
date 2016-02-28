package storage

import (
	"github.com/xh3b4sd/anna/spec"
)

func (rs *redisStorage) GetID() spec.ObjectID {
	rs.Mutex.Lock()
	defer rs.Mutex.Unlock()
	return rs.ID
}

func (rs *redisStorage) GetType() spec.ObjectType {
	rs.Mutex.Lock()
	defer rs.Mutex.Unlock()
	return rs.Type
}
