package factoryserver

import (
	"github.com/xh3b4sd/anna/spec"
)

func (fs *factoryServer) UnmarshalJSON(bytes []byte) error {
	fs.Log.WithTags(spec.Tags{L: "D", O: fs, T: nil, V: 15}, "call UnmarshalJSON")
	return nil
}
