// Package server implements a HTTP server to provide Anna's API over network.
package server

import (
	"net/http"
	"sync"
	"time"

	"github.com/tylerb/graceful"
	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/instrumentation/memory"
	"github.com/xh3b4sd/anna/log"
	logcontrol "github.com/xh3b4sd/anna/server/control/log"
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
func DefaultConfig() Config {
	newInstrumentation, err := memory.NewInstrumentation(memory.DefaultInstrumentationConfig())
	if err != nil {
		panic(err)
	}

	newLogControl, err := logcontrol.NewControl(logcontrol.DefaultControlConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.
		Instrumentation: newInstrumentation,
		Log:             log.NewLog(log.DefaultConfig()),
		LogControl:      newLogControl,
		TextGateway:     gateway.NewGateway(gateway.DefaultConfig()),
		TextInterface:   nil,

		// Settings.
		Addr: "127.0.0.1:9119",
	}

	return newConfig
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
		Server: &graceful.Server{
			NoSignalHandling: true,
			Server: &http.Server{
				Addr: config.Addr,
			},
			Timeout: 3 * time.Second,
		},
		ShutdownOnce: sync.Once{},
		Type:         spec.ObjectType(ObjectTypeServer),
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

	BootOnce     sync.Once
	ID           spec.ObjectID
	Mutex        sync.Mutex
	Server       *graceful.Server
	ShutdownOnce sync.Once
	Type         spec.ObjectType
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
		newTextInterfaceHandlers := text.NewHandlers(ctx, s.TextInterface)
		for url, handler := range newTextInterfaceHandlers {
			http.Handle(url, handler)
		}

		// Server.
		go func() {
			s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 14}, "server starts to listen on '%s'", s.Addr)
			err := s.Server.ListenAndServe()
			if err != nil {
				s.Log.WithTags(spec.Tags{L: "E", O: s, T: nil, V: 4}, "%#v", maskAny(err))
			}
		}()
	})
}

func (s *server) Shutdown() {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call Shutdown")

	s.ShutdownOnce.Do(func() {
		// Stop the server and wait for it to be stopped.
		s.Server.Stop(s.Server.Timeout)
		<-s.Server.StopChan()
	})
}
