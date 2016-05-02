package spec

// CLGIndex represents the CLG index. It provides the CLGCollection. The
// CLGCollection provides all CLGs that can be used.
type CLGIndex interface {
	// Boot initializes and starts the whole CLG index like booting a machine.
	// The call to Boot blocks until the CLG index is completely initialized, so
	// you might want to call it in a separate goroutine.
	Boot()

	CreateCLGProfile(clgCollection CLGCollection, clgName string, canceler <-chan struct{}) (CLGProfile, error)

	// CreateCLGProfiles checks all CLGs optained by the given collection,
	// whether they have proper profiles. In case a proper profile exists, it
	// stays untouched. In case a profile is outdated or not complete, it will be
	// modified. A CLG profile is used to describe the CLGs state, e.g.
	// its interface, functionality and identity.
	CreateCLGProfiles(clgCollection CLGCollection) error

	// GetCLGCollection returns the collection of CLGs the index is responsible
	// for.
	GetCLGCollection() CLGCollection

	// TODO comment
	GetCLGProfileByName(clgName string) (CLGProfile, error)

	Object

	// Shutdown ends all processes of the CLG index like shutting down a machine.
	// The call to Shutdown blocks until the CLG index is completely shut down,
	// so you might want to call it in a separate goroutine.
	Shutdown()

	StoreCLGProfile(clgProfile CLGProfile) error
}
