package clg

import (
	"testing"
	"time"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewIndex(t *testing.T) spec.CLGIndex {
	newIndexConfig, err := DefaultIndexConfig()
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	newIndex, err := NewIndex(newIndexConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newIndex
}

// Test_Index_Index_BootShutdown is a useful test on multiple levels. It
// ensures that Boot and Shutdown work together, even when called multiple
// times. Note on Boot we start a workerpool to generate CLG profiles. Thus a
// lot of things are kicked off during this test. This increases the coverage
// without testing any further specific detail.
func Test_Index_Index_BootShutdown(t *testing.T) {
	newIndex := testMaybeNewIndex(t)

	go newIndex.Boot()
	go newIndex.Boot()

	time.Sleep(2 * time.Second)

	newIndex.Shutdown()
	newIndex.Shutdown()
}
