package profile

import (
	"github.com/xh3b4sd/anna/key"
)

func (g *generator) key(f string, v ...interface{}) string {
	return key.NewSysKey(g, f, v...)
}
