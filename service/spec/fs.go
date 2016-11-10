package spec

import (
	"os"
)

// FS provides certain file system implementations for abstraction and
// testing reasons.
type FS interface {
	Configure() error

	Metadata() map[string]string

	// ReadFile is aligned to ioutil.ReadFile.
	ReadFile(filename string) ([]byte, error)

	Service() Collection

	SetServiceCollection(sc Collection)

	// WriteFile is aligned to ioutil.WriteFile.
	WriteFile(filename string, bytes []byte, perm os.FileMode) error

	Validate() error
}
