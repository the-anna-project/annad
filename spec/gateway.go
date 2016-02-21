package spec

type Listener func(newSignal Signal) (Signal, error)

type Gateway interface {
	Close()

	Object

	// Listen is waiting for signals comming in. Received signals are provided to
	// the Listener. The given closer can be used to end the listening.
	Listen(listener Listener, closer <-chan struct{})

	// Send forwards the given Signal and awaits the corresponding response. Once
	// received, the Signal is returned. The given closer can be used to end the
	// waiting.
	Send(newSignal Signal, closer <-chan struct{}) (Signal, error)
}
