package scheduler

import (
	"github.com/xh3b4sd/anna/spec"
)

func (j *job) GetID() spec.ObjectID {
	return j.ID
}

func (j *job) GetType() spec.ObjectType {
	return j.Type
}
