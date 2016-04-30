package spec

// SmartMap represents a smart implementation of golang's map[string]string.
// This can be used to ease the storage of arbitrary typed data. It is simple
// storing a stringified key value pair (map[string]string). It is rather hard
// to store arbitrary data as provided by something like
// map[string]interface{}. SmartMap implements helper to convert types back and
// forth to ease the converting process.
type SmartMap interface {
	// GetStringMap returns the SmartMap's string map representation.
	GetStringMap() map[string]string

	// GetStringString tries to lookup the given key. In case the given key was
	// found within the smart map, it is returned. Otherwise an error is
	// returned.
	GetStringString(key string) (string, error)

	// GetStringStringSlice tries to lookup the given key. In case the given key
	// was found within the smart map, it is returned. Otherwise an error is
	// returned.
	GetStringStringSlice(key string) ([]string, error)

	// SetStringString tries to set the given key value pair to the smart app.
	SetStringString(key, value string) error

	// SetStringStringSlice tries to set the given key value pair to the smart
	// app.
	SetStringStringSlice(key string, value []string) error
}
