package gateway

import ()

type Gateway interface {
	Close()

	Open()

	ReceiveSignal() Signal

	SendSignal(signal Signal)
}

func NewGateway() Gateway {
	return gateway{
		Link:   make(chan Signal, 1000),
		Closer: make(chan struct{}),
		Opener: make(chan struct{}),
	}
}

type gateway struct {
	Link   chan Signal
	Closer chan struct{}
	Opener chan struct{}
}

func (g gateway) Close() {
	g.Closer <- struct{}{}
}

func (g gateway) Open() {
	g.Opener <- struct{}{}
}

func (g gateway) ReceiveSignal() Signal {
	return <-g.Link
}

func (g gateway) SendSignal(signal Signal) {
	select {
	case <-g.Closer:
		// In case the Closer kicks in, the gateway is closed until the opener
		// receives its signal.
		<-g.Opener
	default:
		g.Link <- signal
	}
}
