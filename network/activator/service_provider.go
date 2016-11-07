package activator

import (
	servicespec "github.com/xh3b4sd/anna/service/spec"
)

func (a *activator) Service() servicespec.Collection {
	return a.ServiceCollection
}
