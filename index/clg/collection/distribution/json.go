package distribution

import (
	"encoding/json"
)

// distributionClone is for making use of the stdlib json implementation. The
// distribution object implements its own marshaler and unmarshaler but only to
// provide json implementations for spec.Distribution. Note, not redirecting
// the type will cause infinite recursion.
type distributionClone distribution

func (d *distribution) MarshalJSON() ([]byte, error) {
	newDistribution := distributionClone(d)

	raw, err := json.Marshal(&newDistribution)
	if err != nil {
		return nil, maskAny(err)
	}

	return raw, nil
}

func (d *distribution) UnmarshalJSON(b []byte) error {
	newDistribution := distributionClone{}

	err := json.Unmarshal(b, &newDistribution)
	if err != nil {
		return maskAny(err)
	}

	d.Name = newDistribution.Name
	d.ID = newDistribution.ID
	d.StaticChannels = newDistribution.StaticChannels
	d.Type = newDistribution.Type
	d.Vectors = newDistribution.Vectors

	return nil
}
