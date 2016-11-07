package spec

// StringMap represents a storage implementation managing operations
// against string maps.
type StringMap interface {
	// GetStringMap returns the hash map stored under the given key.
	GetStringMap(key string) (map[string]string, error)

	// SetStringMap stores the given stringMap under the given key.
	SetStringMap(key string, stringMap map[string]string) error
}
