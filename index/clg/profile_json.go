package clg

import (
	"encoding/json"
)

// clgProfileClone is for making use of the stdlib json implementation. The
// clgProfile object implements its own marshaler and unmarshaler but only to
// provide json implementations for spec.CLGProfile. Note, not redirecting the
// type will cause infinite recursion.
type clgProfileClone clgProfile

func (p *clgProfile) MarshalJSON() ([]byte, error) {
	newCLGProfile := clgProfileClone(*p)

	raw, err := json.Marshal(newCLGProfile)
	if err != nil {
		return nil, maskAny(err)
	}

	return raw, nil
}

func (p *clgProfile) UnmarshalJSON(b []byte) error {
	newCLGProfile := clgProfileClone{}

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
