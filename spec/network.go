package spec

import (
	"reflect"
)

// NetworkPayload represents the request provided to a CLG to ask it to do some
// work.
type NetworkPayload struct {
	// Args represents the arguments intended to be used for the requested CLG
	// execution, or the output values being calculated during the requested CLG
	// execution.
	Args []reflect.Value

	// Destination represents the ID of the CLG that receives the message.
	Destination ObjectID

	// Sources represents the IDs of the CLGs that sent the message.
	Sources []ObjectID
}

// Network provides a neural network based on dynamic and self improving CLG
// execution. The network provides input and output channels. When input is
// received it is injected into the neural communication. The following neural
// activity calculates outputs which are streamed through the output channel
// back to the requestor.
type Network interface {
	// Activate decides if the requested CLG wants to be activated. To make this
	// decision the given input, the CLGs connections and behaviour properties
	// may be considered.
	Activate(clgID ObjectID, inputs []reflect.Value) (bool, error)

	// Boot initializes and starts the whole network like booting a machine. The
	// call to Boot blocks until the network is completely initialized, so you
	// might want to call it in a separate goroutine.
	Boot()

	// Calculate executes the activated CLG and invokes its actual implemented
	// behaviour. This behaviour can be anything. It is up to the CLG.
	Calculate(clgID ObjectID, inputs []reflect.Value) ([]reflect.Value, error)

	// Execute takes care about a CLG's execution. It represents basically a
	// business logic bundle of Acivate, Calculate and Forward.
	Execute(clgID ObjectID, payload NetworkPayload) error

	// Forward is triggered after the CLGs calculation. Here is decided what to
	// do next. Like Activate, it is up to the CLG if it forwards signals to
	// further CLGs. E.g. a CLG may or may not forward its calculated results to
	// one or more CLGs. All this depends on its inputs, calculated outputs, CLG
	// connections and behaviour properties.
	Forward(clgID ObjectID, inputs, outputs []reflect.Value) error

	// Listen makes the network listen on requests from the outside. Here each
	// CLG input channel is managed. This models Listen as kind of cortex in
	// which impulses are dispatched into all possible direction and finally flow
	// back again.
	Listen()

	Object

	// Shutdown ends all processes of the network like shutting down a machine.
	// The call to Shutdown blocks until the network is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()
}
