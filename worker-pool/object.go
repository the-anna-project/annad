package workerpool

import (
	"github.com/xh3b4sd/anna/spec"
)

func (wp *workerPool) GetID() spec.ObjectID {
	return wp.ID
}

func (wp *workerPool) GetType() spec.ObjectType {
	return wp.Type
}
