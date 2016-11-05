package tracker

import (
	"github.com/xh3b4sd/anna/spec"
)

func (t *tracker) GetID() string {
	return t.ID
}

func (t *tracker) GetType() spec.ObjectType {
	return t.Type
}
