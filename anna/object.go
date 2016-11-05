package main

import (
	"github.com/xh3b4sd/anna/spec"
)

func (a *anna) GetID() string {
	return a.ID
}

func (a *anna) GetType() spec.ObjectType {
	return a.Type
}
