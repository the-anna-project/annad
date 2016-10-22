package memory

import (
	"testing"
)

func Test_Storage_Shutdown(t *testing.T) {
	newStorage := MustNew()

	newStorage.Shutdown()
	newStorage.Shutdown()
	newStorage.Shutdown()
}
