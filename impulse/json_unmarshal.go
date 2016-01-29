package impulse

func (i *impulse) UnmarshalJSON(bytes []byte) error {
	i.Log.V(11).Debugf("call Impulse.UnmarshalJSON")

	err := i.GetState().SetStateFromObjectBytes(bytes)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
