package spec

// StorageCollection represents a collection of storage instances. This scopes
// different storage implementations in a simple container, which can easily be
// passed around.
type StorageCollection interface {
	Boot()
	Connection() StorageService
	// Feature represents a feature storage. It is used to store features of
	// information sequences separately. Because of the used key structures and
	// scanning algorithms the feature storage must only be used to store data
	// serving the same conceptual purpose. For instance random features can be
	// retrieved more efficiently when there are only keys belonging to features.
	// Other data structures in here would make the scanning algorithms less
	// efficient.
	Feature() StorageService
	// General represents a general storage. It is used to store general data
	// which is not stored in specialized storage instances.
	General() StorageService
	SetConnection(c StorageService)
	SetFeature(c StorageService)
	SetGeneral(c StorageService)
	// Shutdown ends all processes of the storage collection like shutting down a
	// machine. The call to Shutdown blocks until the storage collection is
	// completely shut down, so you might want to call it in a separate goroutine.
	Shutdown()
}
