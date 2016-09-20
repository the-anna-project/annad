package spec

import (
	"reflect"
)

// NetworkPayload represents the data container carried around within the
// neural network.
type NetworkPayload interface {
	// GetArgs returns the arguments of the current network payload.
	GetArgs() []reflect.Value

	// GetContext returns the Context the current network payload holds as first
	// argument. If no Context can be found, an error is returned.
	GetContext() (Context, error)

	// GetArgs returns the destination of the current network payload.
	GetDestination() ObjectID

	// GetID returns the object ID of the current network payload.
	GetID() ObjectID

	// GetArgs returns the sources of the current network payload.
	GetSources() []ObjectID

	// SetArgs returns the arguments of the current network payload. In case the
	// given arguments are not valid, SetArgs returns an error. See Validate for
	// more information.
	SetArgs(args []reflect.Value) error

	// String returns the concatenated string representations of the currently
	// configured arguments.
	String() string

	// Validate throws an error if the current network payload is not valid. An
	// network payload is not valid if it is empty, or if it does not satisfy the
	// convention of the CLG interface to have a proper Context as first input
	// parameter.
	Validate() error
}

// Network provides a neural network based on dynamic and self improving CLG
// execution. The network provides input and output channels. When input is
// received it is injected into the neural communication. The following neural
// activity calculates outputs which are streamed through the output channel
// back to the requestor.
type Network interface {
	// Activate decides if the requested CLG should be activated. To make this
	// decision the given network payload and formerly received network payloads
	// are considered. CLGs within the neural network are able to join forces to
	// trigger a CLG together, where their own output types do not satisfy the
	// input interface of the requested CLG. For this case some synchronization
	// is required. That means network payloads need to be queued until the
	// requested CLG can be properly executed with the provided inputs. That is
	// why Activate takes a single payload, which represents the currently
	// received one, and a list of network payloads, which represent the formerly
	// received network payloads. Activate tries to find a combination of all
	// known payloads that satisfy the interface of the requested CLG. Note that
	// queue has a maximum length of the number of input arguments of the
	// requested CLG. In case no match can be found the given single payload is
	// added to queue and the oldest payload is removed. In case some match was
	// found the unified network payload matching the CLG's interface is returned
	// so it can be used for the requested CLG execution. Network payloads being
	// matched will be removed from the new queue which is returned as second
	// value.
	Activate(clg CLG, payload NetworkPayload, queue []NetworkPayload) (NetworkPayload, []NetworkPayload, error)

	// Boot initializes and starts the whole network like booting a machine. The
	// call to Boot blocks until the network is completely initialized, so you
	// might want to call it in a separate goroutine.
	Boot()

	// Calculate executes the activated CLG and invokes its actual implemented
	// behaviour. This behaviour can be anything. It is up to the CLG what it
	// does with the provided NetworkPayload.
	Calculate(clg CLG, payload NetworkPayload) (NetworkPayload, error)

	// Forward is triggered after the CLGs calculation. Here is decided what to
	// do next. Like Activate, it is up to the CLG if it forwards signals to
	// further CLGs. E.g. a CLG may or may not forward its calculated results to
	// one or more CLGs. All this depends on the information provided by the
	// given network payload, the CLG's connections and its therefore resulting
	// behaviour properties.
	Forward(clg CLG, payload NetworkPayload) error

	// Listen makes the network listen on requests from the outside. Here each
	// CLG input channel is managed. This way Listen acts as kind of cortex in
	// which signals are dispatched into all possible direction and finally flow
	// back again. Errors during processing of the neural network will be logged
	// to the provided logger.
	Listen()

	Object

	// Shutdown ends all processes of the network like shutting down a machine.
	// The call to Shutdown blocks until the network is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()

	StorageProvider
}
