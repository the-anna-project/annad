// Package server implements a HTTP server to provide Anna's API over network.
package server

import (
	"net/http"
	"sync"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/instrumentation/memory"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/server/control/log"
	"github.com/xh3b4sd/anna/server/interface/text"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeServer represents the object type of the server object. This is
	// used e.g. to register itself to the logger.
	ObjectTypeServer spec.ObjectType = "server"
)

// Config represents the configuration used to create a new server object.
type Config struct {
	// Dependencies.
	Instrumentation spec.Instrumentation
	Log             spec.Log
	LogControl      spec.LogControl
	TextGateway     spec.Gateway
	TextInterface   spec.TextInterface

	// Settings.

	// Addr is the host:port representation based on the golang convention for
	// net.URL and http.ListenAndServe.
	Addr string
}

// DefaultConfig provides a default configuration to create a new server object
// by best effort.
func DefaultConfig() (Config, error) {
	newInstrumentation, err := memory.NewInstrumentation(memory.DefaultInstrumentationConfig())
	if err != nil {
		return Config{}, maskAny(err)
	}

	newConfig := Config{
		// Dependencies.
		Instrumentation: newInstrumentation,
		Log:             log.NewLog(log.DefaultConfig()),
		LogControl:      logcontrol.NewLogControl(logcontrol.DefaultConfig()),
		TextGateway:     gateway.NewGateway(gateway.DefaultConfig()),
		TextInterface:   nil,

		// Settings.
		Addr: "127.0.0.1:9119",
	}

	return newConfig, nil
}

// NewServer creates a new configured server object.
func NewServer(config Config) (spec.Server, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		panic(err)
	}

	newServer := &server{
		Config: config,

		BootOnce: sync.Once{},
		ID:       newID,
		Mutex:    sync.Mutex{},
		Type:     spec.ObjectType(ObjectTypeServer),
	}

	newServer.Log.Register(newServer.GetType())

	if newServer.LogControl == nil {
		return nil, maskAnyf(invalidConfigError, "log control must not be empty")
	}
	if newServer.TextInterface == nil {
		return nil, maskAnyf(invalidConfigError, "text interface must not be empty")
	}

	return newServer, nil
}

type server struct {
	Config

	BootOnce sync.Once
	ID       spec.ObjectID
	Mutex    sync.Mutex
	Type     spec.ObjectType
}

func (s *server) Boot() {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call Boot")

	s.BootOnce.Do(func() {
		ctx := context.Background()

		// Log control.
		newLogControlHandlers := logcontrol.NewHandlers(ctx, s.LogControl)
		for url, handler := range newLogControlHandlers {
			http.Handle(url, handler)
		}

		// Instrumentation.
		http.Handle(s.Instrumentation.GetHTTPEndpoint(), s.Instrumentation.GetHTTPHandler())

		// Text interface.
		newTextInterfaceHandlers := textinterface.NewHandlers(ctx, s.TextInterface)
		for url, handler := range newTextInterfaceHandlers {
			http.Handle(url, handler)
		}

		s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 14}, "server starts to listen on '%s'", s.Addr)
		err := http.ListenAndServe(s.Addr, nil)
		if err != nil {
			s.Log.WithTags(spec.Tags{L: "E", O: s, T: nil, V: 4}, "%#v", maskAny(err))
		}
	})
}
