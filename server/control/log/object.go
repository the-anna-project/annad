package logcontrol

import (
	"github.com/xh3b4sd/anna/spec"
)

func (lc *logControl) GetID() spec.ObjectID {
	return lc.ID
}

func (lc *logControl) GetType() spec.ObjectType {
	return lc.Type
}
