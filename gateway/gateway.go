// Package gateway implements spec.Gateway and provides a in memory
// communication channel between objects. It decouples objects by design using
// signals being send through the gateway. A signal can transport raw bytes or
// arbitrary structures.
package gateway

import (
	"sync"

	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	ObjectTypeGateway spec.ObjectType = "gateway"
)

type Config struct {
	Log spec.Log `json:"-"`
}

func DefaultConfig() Config {
	newConfig := Config{
		Log: log.NewLog(log.DefaultConfig()),
	}

	return newConfig
}

func NewGateway(config Config) spec.Gateway {
	newGateway := &gateway{
		Closed: false,
		Closer: make(chan struct{}, 1),
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Link:   make(chan spec.Signal, 1000),
		Mutex:  sync.Mutex{},
		Type:   spec.ObjectType(ObjectTypeGateway),
	}

	newGateway.Log.Register(newGateway.GetType())

	return newGateway
}

type gateway struct {
	Closed bool
	Closer chan struct{}

	Config

	ID spec.ObjectID

	Link chan spec.Signal

	Mutex sync.Mutex

	Type spec.ObjectType
}

func (g *gateway) Close() {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	if !g.Closed {
		g.Closer <- struct{}{}
		g.Closed = true
	}
}

func (g *gateway) Listen(listener spec.Listener, closer <-chan struct{}) {
	for {
		select {
		case <-closer:
			return
		case <-g.Closer:
			return
		case receivedSignal := <-g.Link:
			newResponder := receivedSignal.GetResponder()

			respondingSignal, err := listener(receivedSignal)
			if err != nil {
				receivedSignal.SetError(maskAny(err))
				newResponder <- receivedSignal
				continue
			}

			newResponder <- respondingSignal
		}
	}
}

func (g *gateway) Send(newSignal spec.Signal, closer <-chan struct{}) (spec.Signal, error) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()

	if g.Closed {
		return nil, maskAny(gatewayClosedError)
	}

	go func() {
		g.Link <- newSignal
	}()

	newResponder := newSignal.GetResponder()

	select {
	case <-closer:
		return nil, maskAny(signalCanceledError)
	case newSignal = <-newResponder:
		if newSignal.GetError() != nil {
			return nil, maskAny(newSignal.GetError())
		}

		return newSignal, nil
	}
}
