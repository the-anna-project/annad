package respnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (rn *respNet) bootObjectTree() {
	rn.Log.WithTags(spec.Tags{L: "D", O: rn, T: nil, V: 13}, "call bootObjectTree")

	rn.EvalNet.Boot()
	rn.ExecNet.Boot()
	rn.PatNet.Boot()
	rn.PredNet.Boot()
	rn.StratNet.Boot()
}
