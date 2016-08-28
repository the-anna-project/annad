// Package memoryfilesystem implements spec.FileSystem and provides an
// in-memory implementation for abstraction and testing reasons.
package memoryfilesystem

import (
	"os"
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectType represents the object type of the memory file system object.
	// This is used e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "memory-file-system"
)

// Config represents the configuration used to create a new memory file system
// object.
type Config struct {
	// Dependencies.
	Log spec.Log
}

// DefaultConfig provides a default configuration to create a new memory file
// system object.
func DefaultConfig() Config {
	newConfig := Config{
		// Dependencies.
		Log: log.NewLog(log.DefaultConfig()),
	}

	return newConfig
}

// NewFileSystem creates a new configured memory file system.
func NewFileSystem(config Config) spec.FileSystem {
	newFileSystem := &memoryFileSystem{
		Config:  config,
		ID:      id.MustNew(),
		Mutex:   sync.Mutex{},
		Storage: map[string][]byte{},
		Type:    ObjectType,
	}

	return newFileSystem
}

type memoryFileSystem struct {
	Config

	ID      spec.ObjectID
	Mutex   sync.Mutex
	Storage map[string][]byte
	Type    spec.ObjectType
}

func (mfs *memoryFileSystem) ReadFile(filename string) ([]byte, error) {
	mfs.Log.WithTags(spec.Tags{C: nil, L: "D", O: mfs, V: 13}, "call ReadFile")

	if bytes, ok := mfs.Storage[filename]; ok {
		return bytes, nil
	}

	pathErr := &os.PathError{
		Op:   "open",
		Path: filename,
		Err:  noSuchFileOrDirectoryError,
	}

	return nil, maskAny(pathErr)
}

func (mfs *memoryFileSystem) WriteFile(filename string, bytes []byte, perm os.FileMode) error {
	mfs.Log.WithTags(spec.Tags{C: nil, L: "D", O: mfs, V: 13}, "call WriteFile")

	mfs.Storage[filename] = bytes
	return nil
}
