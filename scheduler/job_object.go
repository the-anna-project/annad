package scheduler

import (
	"github.com/xh3b4sd/anna/spec"
)

func (j *job) GetID() spec.ObjectID {
	j.Mutex.Lock()
	defer j.Mutex.Unlock()
	return j.ID
}

func (j *job) GetType() spec.ObjectType {
	j.Mutex.Lock()
	defer j.Mutex.Unlock()
	return j.Type
}
