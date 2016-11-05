package mem

import (
	"os"
	"sync"

	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/service/id"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectType represents the object type of the memory file system object.
	// This is used e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "memory-file-system"
)

// ServiceConfig represents the configuration used to create a new memory file
// system service object.
type ServiceConfig struct {
	// Dependencies.
	Log spec.Log
}

// DefaultServiceConfig provides a default configuration to create a new memory
// file system service object.
func DefaultServiceConfig() ServiceConfig {
	newConfig := ServiceConfig{
		// Dependencies.
		Log: log.New(log.DefaultConfig()),
	}

	return newConfig
}

// NewService creates a new configured memory file system.
func NewService(config ServiceConfig) (servicespec.FS, error) {
	newService := &service{
		ServiceConfig: config,
		ID:            id.MustNew(),
		Mutex:         sync.Mutex{},
		Storage:       map[string][]byte{},
		Type:          ObjectType,
	}

	if newService.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}

	return newService, nil
}

// MustNewService creates either a new default configured id service object, or
// panics.
func MustNewService() servicespec.FS {
	newService, err := NewService(DefaultServiceConfig())
	if err != nil {
		panic(err)
	}

	return newService
}

type service struct {
	ServiceConfig

	ID      string
	Mutex   sync.Mutex
	Storage map[string][]byte
	Type    spec.ObjectType
}

func (s *service) ReadFile(filename string) ([]byte, error) {
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call ReadFile")

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
	s.Log.WithTags(spec.Tags{C: nil, L: "D", O: s, V: 13}, "call WriteFile")

	s.Storage[filename] = bytes
	return nil
}
