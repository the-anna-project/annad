// Package id provides a simple ID generating factory using pseudo random
// strings.
package id

import (
	"github.com/xh3b4sd/anna/spec"
)

// MustNew returns a new spec.ObjectID of type Hex128. In case of any error
// this method panics.
func MustNew() spec.ObjectID {
	newID, err := MustNewFactory().WithType(Hex128)
	if err != nil {
		panic(err)
	}

	return newID
}
