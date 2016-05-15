package spec

import (
	"encoding/json"
)

// CLGProfile contains information of a certain CLG.
type CLGProfile interface {
	// Equals checks whether the current CLG profile is equal to the given one.
	Equals(CLGProfile) bool

	// GetBody returns the profile's implemented CLG method body as string
	// representation.
	GetBody() string

	// GetHasChanged returns a bool describing whether a CLG changed.
	GetHasChanged() bool

	// GetHash returns the checksum of the profile's body.
	GetHash() string

	// GetInputs returns the profile's implemented CLG method input parameter
	// types.
	GetInputs() []string

	// GetName returns the name of the CLG this profile is associated with.
	GetName() string

	// GetOutputs returns the profile's implemented CLG method output parameter
	// types.
	GetOutputs() []string

	json.Marshaler

	json.Unmarshaler

	Object

	// SetHashChanged provides a way to set the profile's HasChanged property.
	// See SetID for more background.
	SetHashChanged(hasChanged bool)

	// SetID provides a way to set the ID of a profile. This should only be used
	// during the process of checking profile changes. When checking profile
	// changes on boot a new profile is created and compared to the one that
	// might already exists. The creation is necessary to compare using
	// CLGProfile.Equals. Once a profile is created it already has a new ID.
	// Because we want to obtain IDs of CLG profiles we need to set the ID of the
	// newly created profile to the ID that is already known and used for the
	// profile. That is why this method is necessary.
	SetID(id ObjectID)
}
