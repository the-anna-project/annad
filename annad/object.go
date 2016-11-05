package main

import (
	"github.com/xh3b4sd/anna/spec"
)

func (a *annad) GetID() string {
	return a.ID
}

func (a *annad) GetType() spec.ObjectType {
	return a.Type
}
