package id

import (
	"testing"
)

func Test_IDService_MustNew(t *testing.T) {
	if MustNew() == MustNew() {
		t.Fatal("expected", false, "got", true)
	}
}
