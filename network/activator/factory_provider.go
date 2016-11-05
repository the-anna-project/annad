package activator

import (
	"github.com/xh3b4sd/anna/spec"
)

func (a *activator) Service() spec.ServiceCollection {
	return a.ServiceCollection
}
