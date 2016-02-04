package network

func (n *network) UnmarshalJSON(bytes []byte) error {
	n.Log.V(11).Debugf("call Network.UnmarshalJSON")

	err := n.GetState().SetStateFromObjectBytes(bytes)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
