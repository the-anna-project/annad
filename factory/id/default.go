package id

import (
	"github.com/xh3b4sd/anna/spec"
)

// MustNew returns a new spec.ObjectID of type Hex128. In case of any error
// this method panics.
func MustNew() spec.ObjectID {
	newIDFactory, err := NewFactory(DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(Hex128)
	if err != nil {
		panic(err)
	}

	return newID
}
