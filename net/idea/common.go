package ideanet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (in *ideaNet) bootObjectTree() {
	in.Log.WithTags(spec.Tags{L: "D", O: in, T: nil, V: 13}, "call bootObjectTree")

	in.EvalNet.Boot()
	in.ExecNet.Boot()
	in.PatNet.Boot()
	in.PredNet.Boot()
	in.StratNet.Boot()
}
