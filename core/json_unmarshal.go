package core

func (c *core) UnmarshalJSON(bytes []byte) error {
	c.Log.V(11).Debugf("call Core.UnmarshalJSON")

	err := c.GetState().SetStateFromObjectBytes(bytes)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
