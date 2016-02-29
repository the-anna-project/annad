package spec

// Storage represents a persistency management object.
type Storage interface {
	// Get returns data associated with key. This is a simple key-value
	// relationship.
	Get(key string) (string, error)

	// GetElementsByScore looks up all elements associated with the given score.
	// To limit the number of returned elements, maxElements ca be used.
	GetElementsByScore(key string, score float32, maxElements int) ([]string, error)

	// GetHighestElementScore searches a list that is ordered by their element's
	// score, and returns the element and its corresponding score, where score is
	// the highest within the searched list.
	GetHighestElementScore(key string) (string, float32, error)

	Object

	// Set stores the given key value pair. Once persisted, value can be
	// retrieved using Get.
	Set(key string, value string) error
}
