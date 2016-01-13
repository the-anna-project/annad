package textinterface

import (
	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

var (
	EmptyError = errgo.New("empty")
)
