package server

import (
	"log"
	"net"
	"net/http"
	"strconv"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/server/interface/text"
)

type ServerConfig struct {
	Host        string
	Port        int
	TextGateway gateway.Gateway
}

func DefaultServerConfig() ServerConfig {
	return ServerConfig{
		Host:        "127.0.0.1",
		Port:        9119,
		TextGateway: nil,
	}
}

type Server interface {
	// Listen starts a server based on the given configuration. The call to Boot
	// is blocking, so you might want to call it in a separate goroutine.
	Listen()
}

func NewServer(config ServerConfig) Server {
	return server{
		ServerConfig: config,
	}
}

type server struct {
	ServerConfig
}

func (s server) Listen() {
	ctx := context.Background()

	// text interface
	newTextInterfaceConfig := textinterface.DefaultTextInterfaceConfig()
	newTextInterfaceConfig.TextGateway = s.TextGateway
	newTextInterface := textinterface.NewTextInterface(newTextInterfaceConfig)
	newTextInterfaceHandlers := textinterface.NewHandlers(ctx, newTextInterface)

	// http
	for url, handler := range newTextInterfaceHandlers {
		http.Handle(url, handler)
	}
	log.Fatal(http.ListenAndServe(net.JoinHostPort(s.Host, strconv.Itoa(s.Port)), nil))
}
