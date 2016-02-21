package main

import (
	"github.com/xh3b4sd/anna/spec"
)

func (a *anna) UnmarshalJSON(bytes []byte) error {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 15}, "call UnmarshalJSON")
	return nil
}
