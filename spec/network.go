package spec

import (
	"encoding/json"
	"reflect"
)

// NetworkPayload represents the data container carried around within the
// neural network.
type NetworkPayload interface {
	// GetArgs returns the arguments of the current network payload.
	GetArgs() []reflect.Value

	// GetCLGInput returns a list of arguments intended to be provided as input
	// for a CLG's execution. The list of arguments exists of the arguments
	// configured to the network payload and the context configured to the network
	// payload. Note that the context is always the first argument in the list.
	GetCLGInput() []reflect.Value

	// GetContext returns the context of the current network payload.
	GetContext() Context

	// GetArgs returns the destination of the current network payload, which must
	// be the ID of a CLG registered within the neural network.
	GetDestination() ObjectID

	// GetID returns the object ID of the current network payload.
	GetID() ObjectID

	// GetArgs returns the sources of the current network payload, which must be
	// the ID of a CLG registered within the neural network. One allowed exception
	// is the very first source of the very first network payload, which is
	// created within the network when user input is received to forward it to the
	// input CLG.
	GetSources() []ObjectID

	json.Marshaler
	json.Unmarshaler

	// SetArgs sets the arguments of the current network payload.
	SetArgs(args []reflect.Value)

	// String returns the concatenated string representations of the currently
	// configured arguments.
	String() string

	// Validate throws an error if the current network payload is not valid. An
	// network payload is not valid if it does ot have any context, destination or
	// sources defined.
	Validate() error
}

// Network provides a neural network based on dynamic and self improving CLG
// execution. The network provides input and output channels. When input is
// received it is injected into the neural communication. The following neural
// activity calculates outputs which are streamed through the output channel
// back to the requestor.
type Network interface {
	// Activate decides if a requested CLG shall be activated. It needs to be
	// decided if the requested CLG shall be activated. In case the requested CLG
	// shall be activated it needs to be decided how this activation shall happen.
	// For more information on the activation process see Activator.Activate.
	// This is how it might look like when a CLG is requested and multiple CLGs
	// forward signals to it.
	//
	//     |-----|     |-----|     |-----|     |-----|     |-----|
	//     | CLG |     | CLG |     | CLG |     | CLG |     | CLG |
	//     |-----|     |-----|     |-----|     |-----|     |-----|
	//        |           |           |           |           |
	//        V           V           V           V           V
	//        -------------------------------------------------
	//                                |
	//                                V
	//                             |-----|
	//                             | CLG |
	//                             |-----|
	//
	Activate(CLG CLG, networkPayload NetworkPayload) (NetworkPayload, error)

	// Boot initializes and starts the whole network like booting a machine. The
	// call to Boot blocks until the network is completely initialized, so you
	// might want to call it in a separate goroutine.
	//
	// Boot makes the network listen on requests from the outside. Here each
	// CLG input channel is managed. This way Listen acts as kind of cortex in
	// which signals are dispatched into all possible direction and finally flow
	// back again. Errors during processing of the neural network will be logged
	// to the provided logger.
	Boot()

	// Calculate executes the activated CLG and invokes its actual implemented
	// behaviour. This behaviour can be anything. It is up to the CLG what it
	// does with the provided NetworkPayload.
	Calculate(CLG CLG, networkPayload NetworkPayload) (NetworkPayload, error)

	// EventListener is a worker pool function which is executed multiple times
	// concurrently to listen for network events. A network event is qualified by
	// a network payload being queued and waiting to be processed. A network event
	// and its processing represents an event associated with a very specific CLG.
	// Thus EventListener represents the entrypoint for every single CLG execution
	// step. canceler is provided by the worker pool that executes EventListener.
	// Therefore EventListener should respect canceler by implementing a clean
	// shutdown behaviour. EventListener calls EventHandler.
	EventListener(canceler <-chan struct{}) error

	// EventHandler effectively executes a network event associated to a CLG and
	// the corresponding network payload. EventHandler is called by EventListener.
	EventHandler(CLG CLG, networkPayload NetworkPayload) error

	FactoryProvider

	// Forward is triggered after the CLGs calculation. Here will be decided what
	// to do next. Like Activate, it is up to the CLG if it forwards signals to
	// further CLGs. E.g. a CLG may or may not forward its calculated results to
	// one or more CLGs. All this depends on the information provided by the given
	// network payload, the CLG's connections and its therefore resulting
	// behaviour properties. This is how it might look like when one CLG forwards
	// signals to multiple other CLGs.
	//
	//                             |-----|
	//                             | CLG |
	//                             |-----|
	//                                |
	//                                V
	//        -------------------------------------------------
	//        |           |           |           |           |
	//        V           V           V           V           V
	//     |-----|     |-----|     |-----|     |-----|     |-----|
	//     | CLG |     | CLG |     | CLG |     | CLG |     | CLG |
	//     |-----|     |-----|     |-----|     |-----|     |-----|
	//
	Forward(CLG CLG, networkPayload NetworkPayload) error

	GatewayProvider

	// InputListener is a worker pool function which is executed multiple times
	// concurrently to listen for network inputs. A network input is qualified by
	// information sequences sent by clients who request some calculation from the
	// network. InputListener also calls InputHandler.
	InputListener(canceler <-chan struct{}) error

	// InputHandler effectively executes the network input by invoking the input
	// CLG using the incoming text request. InputHandler is called by
	// InputListener.
	InputHandler(CLG CLG, textRequest TextRequest) error

	Object

	// Shutdown ends all processes of the network like shutting down a machine.
	// The call to Shutdown blocks until the network is completely shut down, so
	// you might want to call it in a separate goroutine.
	Shutdown()

	// Track tracks connections being created to learn from connection path
	// patterns. Various data structures are stored to observe the behaviour of
	// the neural network to act accordingly.
	Track(CLG CLG, networkPayload NetworkPayload) error

	StorageProvider
}
