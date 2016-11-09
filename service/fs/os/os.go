// Package os implements spec.FS and provides a real OS
// bound file system implementation.
package os

import (
	"io/ioutil"
	"os"

	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// Config represents the configuration used to create a new OS file system
// object.
type Config struct {
	// Dependencies.
	ServiceCollection servicespec.Collection
}

// DefaultConfig provides a default configuration to create a new OS file
// system object.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		ServiceCollection: nil,
	}

	return newConfig
}

// New creates a new configured OS file system.
func New(config Config) (servicespec.FS, error) {
	newService := &service{
		Config: config,
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
	newService.Metadata["kind"] = "os"
	newService.Metadata["name"] = "file-system"
	newService.Metadata["type"] = "service"

	return newService
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
}

func (s *service) ReadFile(filename string) ([]byte, error) {
	s.Service().Log().Line("func", "ReadFile")

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, maskAny(err)
	}

	return bytes, nil
}

func (s *service) WriteFile(filename string, bytes []byte, perm os.FileMode) error {
	s.Service().Log().Line("func", "WriteFile")

	err := ioutil.WriteFile(filename, bytes, perm)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
