package outnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (on *outNet) bootObjectTree() {
	on.Log.WithTags(spec.Tags{L: "D", O: on, T: nil, V: 13}, "call bootObjectTree")

	on.EvalNet.Boot()
	on.ExecNet.Boot()
	on.PatNet.Boot()
	on.PredNet.Boot()
	on.StratNet.Boot()
}
