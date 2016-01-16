package gateway_test

import (
	"testing"

	"github.com/xh3b4sd/anna/gateway"
)

// Test_Gateway_001 checks that a signal can be dispatched over a normal
// gateway.
func Test_Gateway_001(t *testing.T) {
	newGateway := gateway.NewGateway()

	input := []byte("hello world")
	newSignal := gateway.NewSignal(input)
	err := newGateway.SendSignal(newSignal)
	if err != nil {
		t.Fatalf("sending signal returned error: %#v", err)
	}

	receivedSignal, err := newGateway.ReceiveSignal()
	if err != nil {
		t.Fatalf("receiving signal returned error: %#v", err)
	}
	if string(receivedSignal.GetBytes()) != string(input) {
		t.Fatalf("received signal '%s' differs from input '%s'", receivedSignal.GetBytes(), input)
	}
}

// Test_Gateway_002 checks that a signal can NOT be dispatched over a closed
// gateway.
func Test_Gateway_002(t *testing.T) {
	newGateway := gateway.NewGateway()
	// Here we close the gateway. Signals should NOT be able to be dispatched.
	newGateway.Close()

	input := []byte("hello world")
	newSignal := gateway.NewSignal(input)
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
	newSignal := gateway.NewSignal(input)
	err := newGateway.SendSignal(newSignal)
	if err != nil {
		t.Fatalf("sending signal returned error: %#v", err)
	}

	receivedSignal, err := newGateway.ReceiveSignal()
	if err != nil {
		t.Fatalf("receiving signal returned error: %#v", err)
	}
	if string(receivedSignal.GetBytes()) != string(input) {
		t.Fatalf("received signal '%s' differs from input '%s'", receivedSignal.GetBytes(), input)
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
		if string(receivedSignal.GetBytes()) != string(input) {
			t.Fatalf("received signal '%s' differs from input '%s'", receivedSignal.GetBytes(), input)
		}
	}()

	<-ready

	newSignal := gateway.NewSignal(input)
	err := newGateway.SendSignal(newSignal)
	if err != nil {
		t.Fatalf("sending signal returned error: %#v", err)
	}
}
