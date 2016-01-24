package spec

type Gateway interface {
	Close()

	Open()

	ReceiveSignal() (Signal, error)

	SendSignal(signal Signal) error
}
