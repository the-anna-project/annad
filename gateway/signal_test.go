package gateway_test

import (
	"fmt"
	"testing"

	"github.com/xh3b4sd/anna/gateway"
)

// Test_Signal_001 checks that setting and getting errors on signals works as
// expected.
func Test_Signal_001(t *testing.T) {
	newSignalConfig := gateway.DefaultSignalConfig()
	newSignal := gateway.NewSignal(newSignalConfig)

	err := newSignal.GetError()
	if err != nil {
		t.Fatalf("Signal.GetError returned error: %#v", err)
	}
	newSignal.SetError(fmt.Errorf("test error"))
	err = newSignal.GetError()
	if err == nil {
		t.Fatalf("Signal.GetError did NOT return error")
	}
	if err.Error() != "test error" {
		t.Fatalf("Signal.GetError returned wrong error")
	}
}

// Test_Signal_002 checks that setting and getting IDs on signals works as
// expected.
func Test_Signal_002(t *testing.T) {
	newSignalConfig := gateway.DefaultSignalConfig()
	newSignalConfig.ID = "testID"
	newSignal := gateway.NewSignal(newSignalConfig)

	ID := newSignal.GetID()
	if ID != "testID" {
		t.Fatalf("Signal.GetID returned wrong error")
	}
}

// Test_Signal_003 checks that setting and getting errors on signals works as
// expected.
func Test_Signal_003(t *testing.T) {
	newSignalConfig := gateway.DefaultSignalConfig()
	newSignal := gateway.NewSignal(newSignalConfig)

	_, err := newSignal.GetObject("foo")
	if !gateway.IsObjectNotFound(err) {
		t.Fatalf("Signal.GetObject did NOT return err")
	}
	newSignal.SetObject("foo", true)
	object, err := newSignal.GetObject("foo")
	if err != nil {
		t.Fatalf("Signal.GetObject returned err: %#v", err)
	}
	if object.(bool) != true {
		t.Fatalf("Signal.GetObject returned invalid object")
	}
}
