package spec

// Storage represents a persistency management object.
type Storage interface {
	// Get returns data associated with key. This is a simple key-value
	// relationship.
	Get(key string) (string, error)

	// GetElementsByScore looks up all elements associated with the given score.
	// To limit the number of returned elements, maxElements ca be used.
	GetElementsByScore(key string, score float32, maxElements int) ([]string, error)

	// GetHighestScoredElements searches a list that is ordered by their
	// element's score, and returns the elements and their corresponding scores,
	// where the highest scored element is the first in the returned list. Note
	// that the list has this scheme.
	//
	//     element1,score1,element2,score2,...
	//
	GetHighestScoredElements(key string, maxElements int) ([]string, error)

	Object

	// Set stores the given key value pair. Once persisted, value can be
	// retrieved using Get.
	Set(key string, value string) error
}
