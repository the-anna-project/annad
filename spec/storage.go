package spec

type Storage interface {
	Get(key string) (string, error)

	GetElementsByScore(key string, score float32, maxElements int) ([]string, error)

	// GetHighestElementScore searches a list that is ordered by their element's
	// score, and returns the element and its corresponding score, where score is
	// the highest within the searched list.
	GetHighestElementScore(key string) (string, float32, error)

	Object

	Set(key string, value string) error
}
