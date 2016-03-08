package main

import (
	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) GetID() spec.ObjectID {
	return a.ID
}

func (a *annactl) GetType() spec.ObjectType {
	return a.Type
}
