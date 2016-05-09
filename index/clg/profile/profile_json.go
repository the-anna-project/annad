package profile

import (
	"encoding/json"
)

// profileClone is for making use of the stdlib json implementation. The
// profile object implements its own marshaler and unmarshaler but only to
// provide json implementations for spec.CLGProfile. Note, not redirecting the
// type will cause infinite recursion.
type profileClone profile

func (p *profile) MarshalJSON() ([]byte, error) {
	newCLGProfile := profileClone(*p)

	raw, err := json.Marshal(newCLGProfile)
	if err != nil {
		return nil, maskAny(err)
	}

	return raw, nil
}

func (p *profile) UnmarshalJSON(b []byte) error {
	newCLGProfile := profileClone{}

	err := json.Unmarshal(b, &newCLGProfile)
	if err != nil {
		return maskAny(err)
	}

	p.Hash = newCLGProfile.Hash
	p.InputTypes = newCLGProfile.InputTypes
	p.InputExamples = newCLGProfile.InputExamples
	p.MethodName = newCLGProfile.MethodName
	p.MethodBody = newCLGProfile.MethodBody
	p.OutputTypes = newCLGProfile.OutputTypes
	p.OutputExamples = newCLGProfile.OutputExamples
	p.RightSideNeighbours = newCLGProfile.RightSideNeighbours

	return nil
}
