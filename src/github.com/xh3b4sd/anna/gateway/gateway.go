package gateway

import (
	"sync"
)

type Gateway interface {
	Close()

	Open()

	ReceiveSignal() (Signal, error)

	SendSignal(signal Signal) error
}

func NewGateway() Gateway {
	g := &gateway{
		Link:   make(chan Signal, 1000),
		Closed: false,
		Mutex:  sync.Mutex{},
	}

	return g
}

type gateway struct {
	Link   chan Signal
	Closed bool
	Mutex  sync.Mutex
}

func (g *gateway) Close() {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	g.Closed = true
}

func (g *gateway) Open() {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	g.Closed = false
}

func (g *gateway) ReceiveSignal() (Signal, error) {
	// Note that we can NOT simply defer the call to Mutex.Unlock, because of the
	// Link channel at the end of this function. Link might block until a signal
	// can be read again. In this case Mutex.Unlock is never called and the mutex
	// causes blocking of Gateway.Close, Gateway.Open and Gateway.SendSignal. So
	// we need to explicitly unlock here.
	g.Mutex.Lock()
	if g.Closed {
		g.Mutex.Unlock()
		return nil, maskAny(gatewayClosedError)
	}
	g.Mutex.Unlock()

	return <-g.Link, nil
}

func (g *gateway) SendSignal(signal Signal) error {
	// Note that we can NOT simply defer the call to Mutex.Unlock, because of the
	// Link channel at the end of this function. Link might block until a signal
	// can be sent again. In this case Mutex.Unlock is never called and the mutex
	// causes blocking of Gateway.Close, Gateway.Open and Gateway.ReceiveSignal.
	// So we need to explicitly unlock here.
	g.Mutex.Lock()
	if g.Closed {
		g.Mutex.Unlock()
		return maskAny(gatewayClosedError)
	}
	g.Mutex.Unlock()

	g.Link <- signal
	return nil
}
