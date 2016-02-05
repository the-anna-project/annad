package filesystemfake

import (
	"os"

	"github.com/xh3b4sd/anna/spec"
)

func NewFileSystem() spec.FileSystem {
	newFileSystem := &fake{
		Storage: map[string][]byte{},
	}

	return newFileSystem
}

type fake struct {
	Storage map[string][]byte
}

func (f *fake) ReadFile(filename string) ([]byte, error) {
	if bytes, ok := f.Storage[filename]; ok {
		return bytes, nil
	}

	pathErr := &os.PathError{
		Op:   "open",
		Path: filename,
		Err:  noSuchFileOrDirectoryError,
	}

	return nil, maskAny(pathErr)
}

func (f *fake) WriteFile(filename string, bytes []byte, perm os.FileMode) error {
	f.Storage[filename] = bytes
	return nil
}
