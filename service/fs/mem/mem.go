// Package mem implements spec.FS and provides an in-memory file system
// implementation for abstraction and testing reasons.
package mem

import (
	"os"
	"sync"

	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// Config represents the configuration used to create a new memory file
// system service object.
type Config struct {
	// Dependencies.
	ServiceCollection servicespec.Collection
}

// DefaultConfig provides a default configuration to create a new memory
// file system service object.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		ServiceCollection: nil,
	}

	return newConfig
}

// New creates a new configured memory file system.
func New(config Config) (servicespec.FS, error) {
	newService := &service{
		Config:  config,
		Mutex:   sync.Mutex{},
		Storage: map[string][]byte{},
	}

	// Dependencies
	if newService.ServiceCollection == nil {
		return nil, maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	id, err := newService.Service().ID().New()
	if err != nil {
		return nil, maskAny(err)
	}
	newService.Metadata["id"] = id
	newService.Metadata["kind"] = "memory"
	newService.Metadata["name"] = "file-system"
	newService.Metadata["type"] = "service"

	return newService, nil
}

// MustNew creates either a new default configured id service, or panics.
func MustNew() servicespec.FS {
	newService, err := New(DefaultConfig())
	if err != nil {
		panic(err)
	}

	return newService
}

type service struct {
	Config

	Metadata map[string]string
	// TODO actually use mutex
	Mutex   sync.Mutex
	Storage map[string][]byte
}

func (s *service) ReadFile(filename string) ([]byte, error) {
	s.Service().Log().Line("func", "ReadFile")

	if bytes, ok := s.Storage[filename]; ok {
		return bytes, nil
	}

	pathErr := &os.PathError{
		Op:   "open",
		Path: filename,
		Err:  noSuchFileOrDirectoryError,
	}

	return nil, maskAny(pathErr)
}

func (s *service) WriteFile(filename string, bytes []byte, perm os.FileMode) error {
	s.Service().Log().Line("func", "WriteFile")

	s.Storage[filename] = bytes
	return nil
}
