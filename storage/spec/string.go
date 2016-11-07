package spec

// String represents a storage implementation managing operations against
// strings.
// TODO the string naming is weird
type String interface {
	// Get returns data associated with key. This is a simple key-value
	// relationship.
	Get(key string) (string, error)

	// GetRandom returns a random key which was formerly stored within the
	// underlying storage.
	GetRandom() (string, error)

	// Remove deletes the given key.
	Remove(key string) error

	// Set stores the given key value pair. Once persisted, value can be
	// retrieved using Get.
	Set(key string, value string) error

	// WalkKeys scans the key space with respect to the given glob and executes
	// the callback for each found key.
	//
	// The walk is throttled. That means some amount of keys are fetched at once
	// from the storage. After all fetched keys are iterated, the next batch of
	// keys is fetched to continue the next iteration, until the whole key space
	// is walked completely. The given closer can be used to end the walk
	// immediately.
	WalkKeys(glob string, closer <-chan struct{}, cb func(key string) error) error
}
