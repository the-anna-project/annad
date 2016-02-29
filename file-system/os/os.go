// Package osfilesystem implementes spec.FileSystem and provides a real OS
// bound implementation.
package osfilesystem

import (
	"io/ioutil"
	builtinos "os"

	"github.com/xh3b4sd/anna/spec"
)

func NewFileSystem() spec.FileSystem {
	newFileSystem := &os{}

	return newFileSystem
}

type os struct{}

func (o *os) ReadFile(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, maskAny(err)
	}

	return bytes, nil
}

func (o *os) WriteFile(filename string, bytes []byte, perm builtinos.FileMode) error {
	err := ioutil.WriteFile(filename, bytes, perm)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
