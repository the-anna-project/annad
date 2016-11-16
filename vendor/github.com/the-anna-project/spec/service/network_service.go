package spec

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// NetworkService provides a neural network based on dynamic and self improving CLG
// execution. The network provides input and output channels. When input is
// received it is injected into the neural communication. The following neural
// activity calculates outputs which are streamed through the output channel
// back to the requestor.
type NetworkService interface {
	// Activate decides if a requested CLG shall be activated. It needs to be
	// decided if the requested CLG shall be activated. In case the requested CLG
	// shall be activated it needs to be decided how this activation shall happen.
	// For more information on the activation process see
	// ActivatorService.Activate. This is how it might look like when a CLG is
	// requested and multiple CLGs forward signals to it.
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
	Activate(clgService CLGService, networkPayload objectspec.NetworkPayload) (objectspec.NetworkPayload, error)
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
	Calculate(clgService CLGService, networkPayload objectspec.NetworkPayload) (objectspec.NetworkPayload, error)
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
	EventHandler(clgService CLGService, networkPayload objectspec.NetworkPayload) error
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
	Forward(clgService CLGService, networkPayload objectspec.NetworkPayload) error
	// InputListener is a worker pool function which is executed multiple times
	// concurrently to listen for network inputs. A network input is qualified by
	// information sequences sent by clients who request some calculation from the
	// network. InputListener also calls InputHandler.
	InputListener(canceler <-chan struct{}) error
	// InputHandler effectively executes the network input by invoking the input
	// CLG using the incoming text request. InputHandler is called by
	// InputListener.
	InputHandler(clgService CLGService, textInput objectspec.TextInput) error
	// Shutdown ends all processes of the network like shutting down a machine.
	// The call to Shutdown blocks until the network is completely shut down, so
	// you might want to call it in a separate goroutine.
	Service() ServiceCollection
	SetServiceCollection(serviceCollection ServiceCollection)
	Shutdown()
	// Track tracks connections being created to learn from connection path
	// patterns. Various data structures are stored to observe the behaviour of
	// the neural network to act accordingly.
	Track(clgService CLGService, networkPayload objectspec.NetworkPayload) error
}
