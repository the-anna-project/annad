package spec

// List represents a storage implementation managing operations against
// lists. Lists here contain arbitrary strings that are ordered by insertion.
type List interface {
	// PopFromList returns the next element from the list identified by the given
	// key. Note that a list is an ordered sequence of arbitrary elements.
	// PushToList and PopFromList are operating according to a "first in, first
	// out" primitive. If the requested list is empty, PopFromList blocks
	// infinitely until an element is added to the list. Returned elements will
	// also be removed from the specified list.
	PopFromList(key string) (string, error)

	// PushToList adds the given element to the list identified by the given key.
	// Note that a list is an ordered sequence of arbitrary elements. PushToList
	// and PopFromList are operating according to a "first in, first out"
	// primitive.
	PushToList(key string, element string) error
}
