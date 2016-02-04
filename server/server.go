package server

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/gateway/spec"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/server/interface/text"
	"github.com/xh3b4sd/anna/spec"
)

type Config struct {
	// Host is the host:port representation based on the golang convention for
	// net.URL and http.ListenAndServe.
	Host string

	Log spec.Log

	TextGateway gatewayspec.Gateway
}

func DefaultConfig() Config {
	newConfig := Config{
		Host:        "127.0.0.1:9119",
		Log:         log.NewLog(log.DefaultConfig()),
		TextGateway: gateway.NewGateway(),
	}

	return newConfig
}

type Server interface {
	// Listen starts a server based on the given configuration. The call to Boot
	// is blocking, so you might want to call it in a separate goroutine.
	Listen()
}

func NewServer(config Config) Server {
	return server{
		Config: config,
	}
}

type server struct {
	Config
}

func (s server) Listen() {
	s.Log.V(11).Debugf("call Server.Listen")

	ctx := context.Background()

	// text interface
	newTextInterfaceConfig := textinterface.DefaultConfig()
	newTextInterfaceConfig.Log = s.Log
	newTextInterfaceConfig.TextGateway = s.TextGateway
	newTextInterface := textinterface.NewTextInterface(newTextInterfaceConfig)
	newTextInterfaceHandlers := textinterface.NewHandlers(ctx, newTextInterface)

	// http
	for url, handler := range newTextInterfaceHandlers {
		http.Handle(url, handler)
	}

	s.Log.V(11).Debugf("server starts to listen on '%s'", s.Host)
	err := http.ListenAndServe(s.Host, nil)
	if err != nil {
		s.Log.V(3).Errorf("%#v", maskAny(err))
	}
}
