// Package server implements a HTTP server to provide Anna's API over network.
package server

import (
	"net/http"
	"sync"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/id"
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
	// dependencies
	Log           spec.Log
	LogControl    spec.LogControl
	TextGateway   spec.Gateway
	TextInterface spec.TextInterface

	// settings

	// Addr is the host:port representation based on the golang convention for
	// net.URL and http.ListenAndServe.
	Addr string
}

// DefaultConfig provides a default configuration to create a new server object
// by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		// dependencies
		Log:           log.NewLog(log.DefaultConfig()),
		LogControl:    logcontrol.NewLogControl(logcontrol.DefaultConfig()),
		TextGateway:   gateway.NewGateway(gateway.DefaultConfig()),
		TextInterface: nil,

		// settings
		Addr: "127.0.0.1:9119",
	}

	return newConfig
}

// NewServer creates a new configured server object.
func NewServer(config Config) (spec.Server, error) {
	newServer := &server{
		Booted: false,
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   spec.ObjectType(ObjectTypeServer),
	}

	newServer.Log.Register(newServer.GetType())

	if newServer.TextInterface == nil {
		return nil, maskAnyf(invalidConfigError, "text interface must not be empty")
	}

	return newServer, nil
}

type server struct {
	Config

	Booted bool
	ID     spec.ObjectID
	Mutex  sync.Mutex
	Type   spec.ObjectType
}

func (s *server) Boot() {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	if s.Booted {
		return
	}
	s.Booted = true

	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call Boot")

	ctx := context.Background()

	// text interface
	newTextInterfaceHandlers := textinterface.NewHandlers(ctx, s.TextInterface)
	for url, handler := range newTextInterfaceHandlers {
		http.Handle(url, handler)
	}

	// log control
	newLogControlHandlers := logcontrol.NewHandlers(ctx, s.LogControl)
	for url, handler := range newLogControlHandlers {
		http.Handle(url, handler)
	}

	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 14}, "server starts to listen on '%s'", s.Addr)
	err := http.ListenAndServe(s.Addr, nil)
	if err != nil {
		s.Log.WithTags(spec.Tags{L: "E", O: s, T: nil, V: 4}, "%#v", maskAny(err))
	}
}
