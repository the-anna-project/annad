// TODO
package filesystemreal

import (
	"io/ioutil"
	"os"

	"github.com/xh3b4sd/anna/spec"
)

func NewFileSystem() spec.FileSystem {
	newFileSystem := &real{}

	return newFileSystem
}

type real struct{}

func (r *real) ReadFile(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, maskAny(err)
	}

	return bytes, nil
}

func (r *real) WriteFile(filename string, bytes []byte, perm os.FileMode) error {
	err := ioutil.WriteFile(filename, bytes, perm)
	if err != nil {
		return maskAny(err)
	}

	return nil
}
