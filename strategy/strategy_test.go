package strategy

import (
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

func testMaybeNewStrategy(t *testing.T) spec.Strategy {
	newConfig := DefaultConfig()
	newConfig.Root = spec.CLG("Sum")
	newStrategy, err := New(newConfig)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	return newStrategy
}

func Test_Strategy_New_Success(t *testing.T) {
	// When this does not panic, the test is successfull.
	testMaybeNewStrategy(t)
}

func Test_Strategy_New_RootError(t *testing.T) {
	newConfig := DefaultConfig()
	// Root configuration is missing.
	newConfig.Root = ""

	_, err := New(newConfig)
	if !IsInvalidConfig(err) {
		t.Fatal("expected", true, "got", false)
	}
}
