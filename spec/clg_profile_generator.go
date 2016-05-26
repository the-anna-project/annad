package spec

// CLGProfileGenerator represents a generator being able to create CLG
// profiles.
type CLGProfileGenerator interface {
	// CreateProfile creates a CLG profile for the given CLG name.
	CreateProfile(clgName string) (CLGProfile, error)

	// GetProfileByName fetches the profile from storage, that is associated with
	// the given profile name.
	GetProfileByName(clgName string) (CLGProfile, error)

	// GetProfileNames returns all profile names to be checked and re/generated.
	GetProfileNames() ([]string, error)

	Object

	// StoreProfile stores the given CLG profile in the configured storage.
	StoreProfile(clgProfile CLGProfile) error
}
