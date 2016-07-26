package id

import (
	"testing"
)

func Test_IDFactory_MustNew(t *testing.T) {
	if MustNew() == MustNew() {
		t.Fatal("expected", false, "got", true)
	}
}
