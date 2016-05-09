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
	newProfile := profileClone(*p)

	raw, err := json.Marshal(newProfile)
	if err != nil {
		return nil, maskAny(err)
	}

	return raw, nil
}

func (p *profile) UnmarshalJSON(b []byte) error {
	newProfile := profileClone{}

	err := json.Unmarshal(b, &newProfile)
	if err != nil {
		return maskAny(err)
	}

	p.Body = newProfile.Body
	p.CreatedAt = newProfile.CreatedAt
	p.Hash = newProfile.Hash
	p.ID = newProfile.ID
	p.Inputs = newProfile.Inputs
	p.Name = newProfile.Name
	p.Outputs = newProfile.Outputs
	p.Type = newProfile.Type

	return nil
}
