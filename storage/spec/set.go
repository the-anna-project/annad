package spec

// Set represents a storage implementation managing operations against
// sets. Sets here contain unique strings that are unordered.
type Set interface {
	// GetAllFromSet returns all elements from the stored stored under key.
	GetAllFromSet(key string) ([]string, error)

	// PushToSet adds the given element to the set identified by the given key.
	// Note that a set is an unordered collection of distinct elements.
	PushToSet(key string, element string) error

	// RemoveFromSet removes the given element from the set identified by the
	// given key.
	RemoveFromSet(key string, element string) error

	// WalkSet scans the set given by key and executes the callback for each found
	// element.
	//
	// The walk is throttled. That means some amount of elements are fetched at
	// once from the storage. After all fetched elements are iterated, the next
	// batch of elements is fetched to continue the next iteration, until the
	// given set is walked completely. The given closer can be used to end the
	// walk immediately.
	WalkSet(key string, closer <-chan struct{}, cb func(element string) error) error
}
