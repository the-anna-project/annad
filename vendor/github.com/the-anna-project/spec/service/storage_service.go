package spec

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// StorageService represents a persistency management object. Different storages may be
// provided in a StorageCollection. Within a receiver function the usage of the
// feature storage may look like this.
//
//     func (s *service) Foo() error {
//       rk, err := s.Storage().Feature().GetRandom()
//       ...
//     }
//
type StorageService interface {
	Boot()

	//
	// List.
	//

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

	//
	// Scored Set.
	//

	// GetElementsByScore looks up all elements associated with the given score.
	// To limit the number of returned elements, maxElements ca be used. Note
	// that the list has this scheme.
	//
	//     element1,element2,...
	//
	GetElementsByScore(key string, score float64, maxElements int) ([]string, error)
	// GetHighestScoredElements searches a list that is ordered by their
	// element's score, and returns the elements and their corresponding scores,
	// where the highest scored element is the first in the returned list. Note
	// that the list has this scheme.
	//
	//     element1,score1,element2,score2,...
	//
	// Note that the resulting list will have the length of maxElements*2,
	// because the list contains the elements and their scores.
	//
	GetHighestScoredElements(key string, maxElements int) ([]string, error)
	// RemoveScoredElement removes the given element from the scored set under
	// key.
	RemoveScoredElement(key string, element string) error
	// SetElementByScore persists the given element in the weighted list
	// identified by key with respect to the given score.
	SetElementByScore(key, element string, score float64) error
	// WalkScoredSet scans the scored set given by key and executes the callback
	// for each found element. Note that the walk might ignores the score order.
	//
	// The walk is throttled. That means some amount of elements are fetched at
	// once from the storage. After all fetched elements are iterated, the next
	// batch of elements is fetched to continue the next iteration, until the
	// given set is walked completely. The given closer can be used to end the
	// walk immediately.
	WalkScoredSet(key string, closer <-chan struct{}, cb func(element string, score float64) error) error

	//
	// Set.
	//

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

	//
	// String Map.
	//

	// GetStringMap returns the hash map stored under the given key.
	GetStringMap(key string) (map[string]string, error)
	// SetStringMap stores the given stringMap under the given key.
	SetStringMap(key string, stringMap map[string]string) error
	// Shutdown ends all processes of the storage like shutting down a machine.
	// The call to Shutdown blocks until the storage is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()

	//
	// String.
	// TODO the string naming is weird, key would fit better
	//

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
	SetAddress(a string)
	SetBackoffFactory(bf func() objectspec.Backoff)
	SetPrefix(p string)
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
}
