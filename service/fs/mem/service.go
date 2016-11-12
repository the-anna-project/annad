// Package mem implements spec.FS and provides an in-memory file system
// implementation for abstraction and testing reasons.
package mem

import (
	"os"
	"sync"

	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// New creates a new memory file system service.
func New() servicespec.FS {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.Collection

	// Settings.

	metadata map[string]string
	// TODO actually use mutex
	mutex   sync.Mutex
	storage map[string][]byte
}

func (s *service) Configure() error {
	// Settings.

	id, err := s.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}
	s.metadata = map[string]string{
		"id":   id,
		"kind": "memory",
		"name": "file-system",
		"type": "service",
	}

	s.mutex = sync.Mutex{}
	s.storage = map[string][]byte{}

	return nil
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) ReadFile(filename string) ([]byte, error) {
	s.Service().Log().Line("func", "ReadFile")

	if bytes, ok := s.storage[filename]; ok {
		return bytes, nil
	}

	pathErr := &os.PathError{
		Op:   "open",
		Path: filename,
		Err:  noSuchFileOrDirectoryError,
	}

	return nil, maskAny(pathErr)
}

func (s *service) Service() servicespec.Collection {
	return s.serviceCollection
}

func (s *service) SetServiceCollection(sc servicespec.Collection) {
	s.serviceCollection = sc
}

func (s *service) WriteFile(filename string, bytes []byte, perm os.FileMode) error {
	s.Service().Log().Line("func", "WriteFile")

	s.storage[filename] = bytes
	return nil
}

func (s *service) Validate() error {
	// Dependencies.
	if s.serviceCollection == nil {
		return maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	return nil
}
