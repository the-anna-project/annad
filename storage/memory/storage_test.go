package memory

import (
	"testing"
)

func Test_Memory_MustNew(t *testing.T) {
	newStorage := MustNew()
	defer newStorage.Shutdown()
}
