package spec

import (
	"os"
)

type FileSystem interface {
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, bytes []byte, perm os.FileMode) error
}
