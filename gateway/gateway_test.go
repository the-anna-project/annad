package gateway

import (
	"fmt"
	"testing"

	"github.com/xh3b4sd/anna/spec"
)

// Test_Gateway_Dispatch checks that a signal can be dispatched back and forth
// over a normal gateway.
func Test_Gateway_Dispatch(t *testing.T) {
	// Prepate test environment.
	newGateway := NewGateway(DefaultConfig())

	input := "hello"
	output := "world"
	var err error

	// Create new signal for the request.
	newSignalConfig := DefaultSignalConfig()
	newSignalConfig.Input = input
	newSignal := NewSignal(newSignalConfig)

	// Receive the signal.
	go func() {
		listener := func(newSignal spec.Signal) (spec.Signal, error) {
			receivedInput := newSignal.GetInput()
			if input != receivedInput.(string) {
				t.Fatalf("received input '%s' differs from original input '%s'", receivedInput, input)
			}

			newSignal.SetOutput(output)

			return newSignal, nil
		}

		newGateway.Listen(listener, nil)
	}()

	// Send the signal.
	newSignal, err = newGateway.Send(newSignal, nil)
	if err != nil {
		t.Fatalf("Gateway.Send did return error: %#v", err)
	}

	// Check received output.
	receivedOutput := newSignal.GetOutput()
	if output != receivedOutput.(string) {
		t.Fatalf("received output '%s' differs from original output '%s'", receivedOutput, output)
	}
}

// Test_Gateway_Dispatch_Gateway_Closed checks that a signal can NOT be
// dispatched back and forth over a closed gateway.
func Test_Gateway_Dispatch_Gateway_Closed(t *testing.T) {
	// Prepate test environment.
	newGateway := NewGateway(DefaultConfig())

	input := "hello"
	output := "world"
	var err error

	// Create new signal for the request.
	newSignalConfig := DefaultSignalConfig()
	newSignalConfig.Input = input
	newSignal := NewSignal(newSignalConfig)

	// Close the gateway.
	newGateway.Close()

	// Receive the signal.
	go func() {
		listener := func(newSignal spec.Signal) (spec.Signal, error) {
			receivedInput := newSignal.GetInput()
			if input != receivedInput.(string) {
				t.Fatalf("received input '%s' differs from original input '%s'", receivedInput, input)
			}

			newSignal.SetOutput(output)

			return newSignal, nil
		}

		newGateway.Listen(listener, nil)
	}()

	// Send the signal.
	receivedSignal, err := newGateway.Send(newSignal, nil)
	if !IsGatewayClosed(err) {
		t.Fatalf("Gateway.Send did NOT return proper error")
	}

	// Check received signal.
	if receivedSignal != nil {
		t.Fatalf("received signal is not nil")
	}

	// Check output of original signal has not changed. The gateway is closed.
	// Applied changes from remote are not allowed to get back to us.
	receivedOutput := newSignal.GetOutput()
	if receivedOutput != nil {
		t.Fatalf("received output is not nil")
	}
}

// Test_Gateway_Dispatch_Listener_Closed checks that a signal can NOT be
// dispatched back and forth over a closed listener.
func Test_Gateway_Dispatch_Listener_Closed(t *testing.T) {
	// Prepate test environment.
	newGateway := NewGateway(DefaultConfig())

	input := "hello"
	output := "world"
	var err error

	// Create new signal for the request.
	newSignalConfig := DefaultSignalConfig()
	newSignalConfig.Input = input
	newSignal := NewSignal(newSignalConfig)

	// Receive the signal.
	listener := func(newSignal spec.Signal) (spec.Signal, error) {
		receivedInput := newSignal.GetInput()
		if input != receivedInput.(string) {
			t.Fatalf("received input '%s' differs from original input '%s'", receivedInput, input)
		}

		newSignal.SetOutput(output)

		return newSignal, nil
	}
	closer := make(chan struct{}, 1)
	closer <- struct{}{}
	newGateway.Listen(listener, closer)

	// Send the signal.
	receivedSignal, err := newGateway.Send(newSignal, nil)
	if !IsGatewayClosed(err) {
		t.Fatalf("Gateway.Send did NOT return proper error")
	}

	// Check received signal.
	if receivedSignal != nil {
		t.Fatalf("received signal is not nil")
	}

	// Check output of original signal has not changed. The gateway is closed.
	// Applied changes from remote are not allowed to get back to us.
	receivedOutput := newSignal.GetOutput()
	if receivedOutput != nil {
		t.Fatalf("received output is not nil")
	}
}

// Test_Gateway_003 checks that error handling during signal dispatching works
// as expected.
func Test_Gateway_003(t *testing.T) {
	// Prepate test environment.
	newGateway := NewGateway(DefaultConfig())

	// Create new signal for the request.
	newSignalConfig := DefaultSignalConfig()
	newSignal := NewSignal(newSignalConfig)

	// Receive the signal.
	go func() {
		listener := func(newSignal spec.Signal) (spec.Signal, error) {
			return nil, fmt.Errorf("test error")
		}

		newGateway.Listen(listener, nil)
	}()

	// Send the signal.
	newSignal, err := newGateway.Send(newSignal, nil)
	if err == nil && err.Error() != "test error" {
		t.Fatalf("Gateway.Send did NOT return proper error")
	}

	// Check received signal.
	if newSignal != nil {
		t.Fatalf("received signal is not nil")
	}
}

// Test_Gateway_004 checks that a signal can NOT be dispatched over a closed
// gateway.
func Test_Gateway_004(t *testing.T) {
	// Prepate test environment.
	newGateway := NewGateway(DefaultConfig())

	input := "hello"
	var err error

	// Create new signal for the request.
	newSignalConfig := DefaultSignalConfig()
	newSignalConfig.Input = input
	newSignal := NewSignal(newSignalConfig)

	// Close the gateway.
	newGateway.Close()

	// Send the signal.
	newSignal, err = newGateway.Send(newSignal, nil)
	if !IsGatewayClosed(err) {
		t.Fatalf("Gateway.Send did NOT return proper error")
	}

	// Check received signal.
	if newSignal != nil {
		t.Fatalf("received signal is not nil")
	}
}

// Test_Gateway_005 checks that a signal can be canceled.
func Test_Gateway_005(t *testing.T) {
	// Prepate test environment.
	newGateway := NewGateway(DefaultConfig())

	input := "hello"
	var err error

	// Create new signal for the request.
	newSignalConfig := DefaultSignalConfig()
	newSignalConfig.Input = input
	newSignal := NewSignal(newSignalConfig)

	// Cancel the signal.
	closer := make(chan struct{}, 1)
	closer <- struct{}{}

	// Send the signal.
	_, err = newGateway.Send(newSignal, closer)
	if !IsSignalCanceled(err) {
		t.Fatalf("Gateway.Send did NOT return proper error")
	}
}
