package spec

// Listener must be provided to Gateway.Listen and is called once a signal is
// received.
type Listener func(newSignal Signal) (Signal, error)

// Gateway provides an in-memor channel to send information back and forth in a
// decoupled manner.
type Gateway interface {
	// Close closes the gateway and thus prevents further usage of it. No
	// listening and sending is possible anymore after closing. Signals send but
	// not yet received while closing get lost.
	Close()

	Object

	// Listen is waiting for signals coming in. Received signals are provided to
	// the Listener. The given closer can be used to end the listening.
	Listen(listener Listener, closer <-chan struct{})

	// Send forwards the given Signal and awaits the corresponding response. Once
	// received, the Signal is returned. The given closer can be used to end the
	// waiting.
	Send(newSignal Signal, closer <-chan struct{}) (Signal, error)
}
