package factoryserver

import (
	"github.com/xh3b4sd/anna/spec"
)

func (fs *factoryServer) MarshalJSON() ([]byte, error) {
	fs.Log.WithTags(spec.Tags{L: "D", O: fs, T: nil, V: 15}, "call MarshalJSON")
	return nil, nil
}
