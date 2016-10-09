package forwarder

import (
	"github.com/xh3b4sd/anna/spec"
)

func (f *forwarder) GetID() spec.ObjectID {
	return f.ID
}

func (f *forwarder) GetType() spec.ObjectType {
	return f.Type
}
