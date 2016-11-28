package service

import (
	"os"
)

// FSService provides certain file system implementations for abstraction and
// testing reasons.
type FSService interface {
	Boot()
	Metadata() map[string]string
	// ReadFile is aligned to ioutil.ReadFile.
	ReadFile(filename string) ([]byte, error)
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
	// WriteFile is aligned to ioutil.WriteFile.
	WriteFile(filename string, bytes []byte, perm os.FileMode) error
}
