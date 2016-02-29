// Package textinterface implements spec.TextInterface and provides a way to
// feed neural networks with text input. To make Anna consume text, there is
// the text interface implemented through the network API.
package textinterface

import (
	"sync"

	"time"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	ObjectTypeTextInterface spec.ObjectType = "text-interface"
)

type Config struct {
	Log spec.Log

	TextGateway spec.Gateway
}

func DefaultConfig() Config {
	return Config{
		Log:         log.NewLog(log.DefaultConfig()),
		TextGateway: gateway.NewGateway(gateway.DefaultConfig()),
	}
}

func NewTextInterface(config Config) spec.TextInterface {
	newInterface := &textInterface{
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   spec.ObjectType(ObjectTypeTextInterface),
	}

	newInterface.Log.Register(newInterface.GetType())

	return newInterface
}

type textInterface struct {
	Config

	ID spec.ObjectID

	Mutex sync.Mutex

	Type spec.ObjectType
}

// TODO this should actually fetch a url from the web
func (ti *textInterface) FetchURL(url string) ([]byte, error) {
	return nil, nil
}

// TODO this should actually read a file from file system
func (ti *textInterface) ReadFile(file string) ([]byte, error) {
	return nil, nil
}

// TODO this should actually be streamed
func (ti *textInterface) ReadStream(stream string) ([]byte, error) {
	return nil, nil
}

// return response
func (ti *textInterface) ReadPlainWithID(ctx context.Context, ID string) (string, error) {
	ti.Log.WithTags(spec.Tags{L: "D", O: ti, T: nil, V: 13}, "call ReadPlainWithID")

	newConfig := gateway.DefaultSignalConfig()
	newSignal := gateway.NewSignal(newConfig)
	newSignal.SetID(ID)

	response, err := ti.waitForSignal(ctx, newSignal)
	if err != nil {
		return "", maskAny(err)
	}

	return response, nil
}

// return ID
func (ti *textInterface) ReadPlainWithPlain(ctx context.Context, plain string) (string, error) {
	ti.Log.WithTags(spec.Tags{L: "D", O: ti, T: nil, V: 13}, "call ReadPlainWithPlain")

	newConfig := gateway.DefaultSignalConfig()
	newSignal := gateway.NewSignal(newConfig)
	newSignal.SetInput(plain)

	response, err := ti.waitForSignal(ctx, newSignal)
	if err != nil {
		return "", maskAny(err)
	}

	return response, nil
}

func (ti *textInterface) waitForSignal(ctx context.Context, newSignal spec.Signal) (string, error) {
	ti.Log.WithTags(spec.Tags{L: "D", O: ti, T: nil, V: 13}, "call waitForSignal")

	var err error

	for {
		newSignal, err = ti.TextGateway.Send(newSignal, ctx.Done())
		if err != nil {
			return "", maskAny(err)
		}

		output := newSignal.GetOutput()
		o := output.(string)

		if o == "" {
			time.Sleep(1 * time.Second)
		} else {
			return o, nil
		}
	}
}
