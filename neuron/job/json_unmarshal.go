package jobneuron

import (
	"github.com/xh3b4sd/anna/spec"
)

func (n *neuron) UnmarshalJSON(bytes []byte) error {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 13}, "call UnmarshalJSON")

	err := n.GetState().SetStateFromObjectBytes(bytes)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
