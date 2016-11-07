package tracker

import (
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

func (t *tracker) Service() servicespec.Collection {
	return t.ServiceCollection
}
