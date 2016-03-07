package gateway

import (
	"fmt"
	"testing"
)

// Test_Signal_001 checks that setting and getting errors on signals works as
// expected.
func Test_Signal_001(t *testing.T) {
	newSignalConfig := DefaultSignalConfig()
	newSignal := NewSignal(newSignalConfig)

	err := newSignal.GetError()
	if err != nil {
		t.Fatalf("Signal.GetError returned error: %#v", err)
	}
	newSignal.SetError(fmt.Errorf("test error"))
	err = newSignal.GetError()
	if err == nil {
		t.Fatalf("Signal.GetError did NOT return proper error")
	}
	if err.Error() != "test error" {
		t.Fatalf("Signal.GetError returned wrong error")
	}
}

func Test_Signal_GetID(t *testing.T) {
	newSignalConfig := DefaultSignalConfig()
	newSignalConfig.ID = "testID"
	newSignal := NewSignal(newSignalConfig)

	ID := newSignal.GetID()
	if ID != "testID" {
		t.Fatal("expected", "testID", "got", ID)
	}
}

// Test_Signal_004 checks that setting and getting output on signals works as
// expected.
func Test_Signal_004(t *testing.T) {
	newSignalConfig := DefaultSignalConfig()
	newSignal := NewSignal(newSignalConfig)

	originalOutput := "foo"

	output := newSignal.GetOutput()
	if output != nil {
		t.Fatalf("output of new signal is not nil")
	}
	newSignal.SetOutput(originalOutput)
	receivedOutput := newSignal.GetOutput()
	if receivedOutput.(string) != "foo" {
		t.Fatalf("received output '%s' of signal differs from original output %s'", receivedOutput, originalOutput)
	}
}
