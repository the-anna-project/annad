// Package readinformationid implements spec.CLG and provides functionality to
// read the information sequence stored under a specific information ID.
package readinformationid

import (
	"fmt"

	"github.com/the-anna-project/spec/object"
)

// calculate fetches the information sequence stored under a specific
// information ID. The information ID is provided by the given context.
func (s *service) calculate(ctx spec.Context) (string, error) {
	informationID, ok := ctx.GetInformationID()
	if !ok {
		return "", maskAnyf(invalidInformationIDError, "must not be empty")
	}

	informationSequenceKey := fmt.Sprintf("information-id:%s:information-sequence", informationID)
	informationSequence, err := s.Service().Storage().General().Get(informationSequenceKey)
	if err != nil {
		return "", maskAny(err)
	}

	return informationSequence, nil
}
