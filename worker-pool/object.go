package workerpool

import (
	"github.com/xh3b4sd/anna/spec"
)

func (wp *workerPool) GetID() spec.ObjectID {
	wp.Mutex.Lock()
	defer wp.Mutex.Unlock()
	return wp.ID
}

func (wp *workerPool) GetType() spec.ObjectType {
	wp.Mutex.Lock()
	defer wp.Mutex.Unlock()
	return wp.Type
}
