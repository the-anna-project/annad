// Package input implements spec.CLG and provides the entry to the neural
// network.
package input

import (
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
)

// calculate fetches the information ID associated with the given information
// sequence. In case the information sequence is not found within the underlying
// storage, a new information ID is generated and used to store the given
// information sequence. In any case the information ID is added to the given
// context.
func (c *clg) calculate(ctx spec.Context, informationSequence string) error {
	informationIDKey := key.NewCLGKey("information-sequence:%s:information-id", informationSequence)
	informationID, err := c.Storage.Get(informationIDKey)
	if storage.IsNotFound(err) {
		// The given information sequence was never seen before. Thus we register it
		// now with its own very unique information ID.
		newID, err := c.IDFactory.New()
		if err != nil {
			return maskAny(err)
		}
		informationID = string(newID)

		// TODO test
		informationSequenceKey := key.NewCLGKey("information-id:%s:information-sequence", informationID)
		err = c.Storage.Set(informationSequenceKey, informationSequence)
		if err != nil {
			return maskAny(err)
		}

		err = c.Storage.Set(informationIDKey, informationID)
		if err != nil {
			return maskAny(err)
		}
	} else if err != nil {
		return maskAny(err)
	}

	ctx.SetInformationID(informationID)

	return nil
}
