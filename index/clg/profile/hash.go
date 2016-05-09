package profile

import (
	"crypto/md5"
	"fmt"

	"github.com/xh3b4sd/anna/spec"
)

func (g *generator) CreateHash(clgBody string) (string, error) {
	g.Log.WithTags(spec.Tags{L: "D", O: g, T: nil, V: 13}, "call CreateHash")

	return fmt.Sprintf("%x\n", md5.Sum([]byte(clgBody))), nil
}
