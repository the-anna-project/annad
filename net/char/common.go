package charnet

import (
	"github.com/xh3b4sd/anna/spec"
)

func (cn *charNet) bootObjectTree() {
	cn.Log.WithTags(spec.Tags{L: "D", O: cn, T: nil, V: 13}, "call bootObjectTree")

	cn.EvalNet.Boot()
	cn.ExecNet.Boot()
	cn.PatNet.Boot()
	cn.PredNet.Boot()
	cn.StratNet.Boot()
}
