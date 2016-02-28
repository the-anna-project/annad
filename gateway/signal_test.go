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

// Test_Signal_002 checks that setting and getting IDs on signals works as
// expected.
func Test_Signal_002(t *testing.T) {
	newSignalConfig := DefaultSignalConfig()
	newSignalConfig.ID = "testID"
	newSignal := NewSignal(newSignalConfig)

	ID := newSignal.GetID()
	if ID != "testID" {
		t.Fatalf("Signal.GetID returned wrong ID")
	}

	// Now using the setter.
	newSignal.SetID("other")
	ID = newSignal.GetID()
	if ID != "other" {
		t.Fatalf("Signal.GetID returned wrong ID")
	}
}

// Test_Signal_003 checks that setting and getting input on signals works as
// expected.
func Test_Signal_003(t *testing.T) {
	newSignalConfig := DefaultSignalConfig()
	newSignal := NewSignal(newSignalConfig)

	originalInput := "foo"

	input := newSignal.GetInput()
	if input != nil {
		t.Fatalf("input of new signal is not nil")
	}
	newSignal.SetInput(originalInput)
	receivedInput := newSignal.GetInput()
	if receivedInput.(string) != "foo" {
		t.Fatalf("received input '%s' of signal differs from original input %s'", receivedInput, originalInput)
	}
}

// Test_Signal_004 checks that setting and getting ouput on signals works as
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
