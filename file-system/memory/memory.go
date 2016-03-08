// Package memoryfilesystem implementes spec.FileSystem and provides an
// in-memory implementation for abstraction and testing reasons.
package memoryfilesystem

import (
	"os"

	"github.com/xh3b4sd/anna/spec"
)

// NewFileSystem creates a new configured memory file system.
func NewFileSystem() spec.FileSystem {
	newFileSystem := &memory{
		Storage: map[string][]byte{},
	}

	return newFileSystem
}

type memory struct {
	Storage map[string][]byte
}

func (m *memory) ReadFile(filename string) ([]byte, error) {
	if bytes, ok := m.Storage[filename]; ok {
		return bytes, nil
	}

	pathErr := &os.PathError{
		Op:   "open",
		Path: filename,
		Err:  noSuchFileOrDirectoryError,
	}

	return nil, maskAny(pathErr)
}

func (m *memory) WriteFile(filename string, bytes []byte, perm os.FileMode) error {
	m.Storage[filename] = bytes
	return nil
}
