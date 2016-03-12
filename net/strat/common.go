package stratnet

import (
	"github.com/xh3b4sd/anna/key"
)

func (sn *stratNet) key(f string, v ...interface{}) string {
	return key.NewNetKey(sn, f, v...)
}
