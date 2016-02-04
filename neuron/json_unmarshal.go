package neuron

//
// character neuron
//

func (cn *characterNeuron) UnmarshalJSON(bytes []byte) error {
	cn.Log.V(11).Debugf("call CharacterNeuron.UnmarshalJSON")

	err := cn.GetState().SetStateFromObjectBytes(bytes)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

//
// first neuron
//

func (fn *firstNeuron) UnmarshalJSON(bytes []byte) error {
	fn.Log.V(11).Debugf("call FirstNeuron.UnmarshalJSON")

	err := fn.GetState().SetStateFromObjectBytes(bytes)
	if err != nil {
		return maskAny(err)
	}

	return nil
}

//
// job neuron
//

func (jn *jobNeuron) UnmarshalJSON(bytes []byte) error {
	jn.Log.V(11).Debugf("call JobNeuron.UnmarshalJSON")

	err := jn.GetState().SetStateFromObjectBytes(bytes)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
