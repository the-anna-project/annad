package gateway_test

import (
	"testing"

	"github.com/xh3b4sd/anna/gateway"
)

// Test_Gateway_001 checks that a signal can be dispatched back and forth over
// a normal gateway.
func Test_Gateway_001(t *testing.T) {
	// Prepate test environment.
	newGateway := gateway.NewGateway()

	input := []byte("hello")
	output := []byte("world")

	// Create new signal for the request.
	newSignalConfig := gateway.DefaultSignalConfig()
	newSignalConfig.Bytes["request"] = input
	newSignal := gateway.NewSignal(newSignalConfig)
	err := newGateway.SendSignal(newSignal)
	if err != nil {
		t.Fatalf("sending signal returned error: %#v", err)
	}

	// Receive signal of the request.
	receivedSignal, err := newGateway.ReceiveSignal()
	if err != nil {
		t.Fatalf("receiving signal returned error: %#v", err)
	}
	request, err := receivedSignal.GetBytes("request")
	if err != nil {
		t.Fatalf("getting bytes returned error: %#v", err)
	}
	if string(request) != string(input) {
		t.Fatalf("received signal bytes '%s' differ from input '%s'", request, input)
	}

	// Fetch the responder and send a response over it.
	responder, err := receivedSignal.GetResponder()
	if err != nil {
		t.Fatalf("getting responder returned error: %#v", err)
	}
	receivedSignal.SetBytes("response", output)
	responder <- receivedSignal

	// Receive the response.
	responder, err = receivedSignal.GetResponder()
	if err != nil {
		t.Fatalf("getting responder returned error: %#v", err)
	}
	receivedSignal = <-responder
	response, err := receivedSignal.GetBytes("response")
	if err != nil {
		t.Fatalf("getting bytes returned error: %#v", err)
	}
	if string(response) != string(output) {
		t.Fatalf("received signal bytes '%s' differ from output '%s'", response, output)
	}
}

// Test_Gateway_002 checks that a signal can NOT be dispatched over a closed
// gateway.
func Test_Gateway_002(t *testing.T) {
	newGateway := gateway.NewGateway()
	// Here we close the gateway. Signals should NOT be able to be dispatched.
	newGateway.Close()

	input := []byte("hello world")

	newSignalConfig := gateway.DefaultSignalConfig()
	newSignalConfig.Bytes["request"] = input
	newSignal := gateway.NewSignal(newSignalConfig)
	err := newGateway.SendSignal(newSignal)
	if err == nil {
		t.Fatalf("sending signal NOT returned error")
	}
	if !gateway.IsGatewayClosed(err) {
		t.Fatalf("sending signal NOT returned proper error")
	}

	receivedSignal, err := newGateway.ReceiveSignal()
	if err == nil {
		t.Fatalf("receiving signal NOT returned error")
	}
	if !gateway.IsGatewayClosed(err) {
		t.Fatalf("sending signal NOT returned proper error")
	}
	if receivedSignal != nil {
		t.Fatalf("receiving signal is NOT nil")
	}
}

// Test_Gateway_003 checks that a signal can be dispatched over a closed
// gateway that was opened again.
func Test_Gateway_003(t *testing.T) {
	newGateway := gateway.NewGateway()
	// Here we close the gateway and open it again. Signals should be able to be
	// dispatched.
	newGateway.Close()
	newGateway.Open()

	input := []byte("hello world")

	newSignalConfig := gateway.DefaultSignalConfig()
	newSignalConfig.Bytes["request"] = input
	newSignal := gateway.NewSignal(newSignalConfig)
	err := newGateway.SendSignal(newSignal)
	if err != nil {
		t.Fatalf("sending signal returned error: %#v", err)
	}

	receivedSignal, err := newGateway.ReceiveSignal()
	if err != nil {
		t.Fatalf("receiving signal returned error: %#v", err)
	}
	request, err := receivedSignal.GetBytes("request")
	if err != nil {
		t.Fatalf("getting bytes returned error: %#v", err)
	}
	if string(request) != string(input) {
		t.Fatalf("received signal bytes '%s' differ from input '%s'", request, input)
	}
}

// Test_Gateway_004 checks that a signal can be dispatched over a normal
// gateway, even when Gateway.ReceiveSignal is called first. This ensures the
// locking mutex works as expected without creating a deadlock for
// Gateway.SendSignal.
func Test_Gateway_004(t *testing.T) {
	newGateway := gateway.NewGateway()

	input := []byte("hello world")

	ready := make(chan struct{})
	go func() {
		ready <- struct{}{}
		receivedSignal, err := newGateway.ReceiveSignal()
		if err != nil {
			t.Fatalf("receiving signal returned error: %#v", err)
		}
		request, err := receivedSignal.GetBytes("request")
		if err != nil {
			t.Fatalf("getting bytes returned error: %#v", err)
		}
		if string(request) != string(input) {
			t.Fatalf("received signal bytes '%s' differ from input '%s'", request, input)
		}
	}()

	<-ready

	newSignalConfig := gateway.DefaultSignalConfig()
	newSignalConfig.Bytes["request"] = input
	newSignal := gateway.NewSignal(newSignalConfig)
	err := newGateway.SendSignal(newSignal)
	if err != nil {
		t.Fatalf("sending signal returned error: %#v", err)
	}
}

// Test_Gateway_005 checks that a signal can be canceled.
func Test_Gateway_005(t *testing.T) {
	// Prepate test environment.
	newGateway := gateway.NewGateway()

	input := []byte("hello")

	// Create new signal for the request.
	newSignalConfig := gateway.DefaultSignalConfig()
	newSignalConfig.Bytes["request"] = input
	newSignal := gateway.NewSignal(newSignalConfig)
	err := newGateway.SendSignal(newSignal)
	if err != nil {
		t.Fatalf("sending signal returned error: %#v", err)
	}

	// Receive signal of the request.
	receivedSignal, err := newGateway.ReceiveSignal()
	if err != nil {
		t.Fatalf("receiving signal returned error: %#v", err)
	}
	request, err := receivedSignal.GetBytes("request")
	if err != nil {
		t.Fatalf("getting bytes returned error: %#v", err)
	}
	if string(request) != string(input) {
		t.Fatalf("received signal bytes '%s' differ from input '%s'", request, input)
	}

	// Actually cancel the signal.
	newSignal.Cancel()

	// Try to fetch the responder. This should not be possible.
	responder, err := receivedSignal.GetResponder()
	if !gateway.IsSignalCanceled(err) {
		t.Fatalf("signal NOT canceled")
	}
	if responder != nil {
		t.Fatalf("responder is NOT nil")
	}
}
