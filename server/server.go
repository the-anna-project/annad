// Package server implements a HTTP server to provide Anna's API over network.
package server

import (
	"net/http"
	"sync"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/server/control/log"
	"github.com/xh3b4sd/anna/server/interface/text"
	"github.com/xh3b4sd/anna/spec"
)

type Config struct {
	// Host is the host:port representation based on the golang convention for
	// net.URL and http.ListenAndServe.
	Host string

	Log spec.Log

	TextGateway spec.Gateway
}

func DefaultConfig() Config {
	newConfig := Config{
		Host:        "127.0.0.1:9119",
		Log:         log.NewLog(log.DefaultConfig()),
		TextGateway: gateway.NewGateway(gateway.DefaultConfig()),
	}

	return newConfig
}

func NewServer(config Config) spec.Server {
	return &server{
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   spec.ObjectType(common.ObjectType.Log),
	}
}

type server struct {
	Config

	ID spec.ObjectID

	Mutex sync.Mutex

	Type spec.ObjectType
}

func (s *server) Boot() {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call Boot")

	ctx := context.Background()

	// text interface
	newTextInterfaceConfig := textinterface.DefaultConfig()
	newTextInterfaceConfig.Log = s.Log
	newTextInterfaceConfig.TextGateway = s.TextGateway
	newTextInterface := textinterface.NewTextInterface(newTextInterfaceConfig)
	newTextInterfaceHandlers := textinterface.NewHandlers(ctx, newTextInterface)
	for url, handler := range newTextInterfaceHandlers {
		http.Handle(url, handler)
	}

	// log control
	newLogControlConfig := logcontrol.DefaultConfig()
	newLogControlConfig.Log = s.Log
	newLogControl := logcontrol.NewLogControl(newLogControlConfig)
	newLogControlHandlers := logcontrol.NewHandlers(ctx, newLogControl)
	for url, handler := range newLogControlHandlers {
		http.Handle(url, handler)
	}

	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 14}, "server starts to listen on '%s'", s.Host)
	err := http.ListenAndServe(s.Host, nil)
	if err != nil {
		s.Log.WithTags(spec.Tags{L: "E", O: s, T: nil, V: 4}, "%#v", maskAny(err))
	}
}
