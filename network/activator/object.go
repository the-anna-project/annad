package activator

import (
	"github.com/xh3b4sd/anna/spec"
)

func (a *activator) GetID() spec.ObjectID {
	return a.ID
}

func (a *activator) GetType() spec.ObjectType {
	return a.Type
}
