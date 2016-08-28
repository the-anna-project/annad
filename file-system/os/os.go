// Package osfilesystem implements spec.FileSystem and provides a real OS
// bound file system implementation.
package osfilesystem

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectType represents the object type of the OS file system object. This
	// is used e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "os-file-system"
)

// Config represents the configuration used to create a new OS file system
// object.
type Config struct {
	// Dependencies.
	Log spec.Log
}

// DefaultConfig provides a default configuration to create a new OS file
// system object.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log: log.NewLog(log.DefaultConfig()),
	}

	return newConfig
}

// NewFileSystem creates a new configured real OS file system.
func NewFileSystem(config Config) spec.FileSystem {
	newFileSystem := &osFileSystem{
		Config: config,
		ID:     id.MustNew(),
		Mutex:  sync.Mutex{},
		Type:   ObjectType,
	}

	newFileSystem.Log.Register(newFileSystem.GetType())

	return newFileSystem
}

type osFileSystem struct {
	Config

	ID    spec.ObjectID
	Mutex sync.Mutex
	Type  spec.ObjectType
}

func (osfs *osFileSystem) ReadFile(filename string) ([]byte, error) {
	osfs.Log.WithTags(spec.Tags{C: nil, L: "D", O: osfs, V: 13}, "call ReadFile")

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, maskAny(err)
	}

	return bytes, nil
}

func (osfs *osFileSystem) WriteFile(filename string, bytes []byte, perm os.FileMode) error {
	osfs.Log.WithTags(spec.Tags{C: nil, L: "D", O: osfs, V: 13}, "call WriteFile")

	err := ioutil.WriteFile(filename, bytes, perm)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
