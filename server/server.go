package server

import (
	"fmt"
	"net"
	"net/http"
	"strconv"

	"golang.org/x/net/context"

	gatewayspec "github.com/xh3b4sd/anna/gateway/spec"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/server/interface/text"
	"github.com/xh3b4sd/anna/spec"
)

type Config struct {
	Host string

	Log spec.Log `json:"-"`

	Port        int
	TextGateway gatewayspec.Gateway
}

func DefaultConfig() Config {
	newConfig := Config{
		Host:        "127.0.0.1",
		Log:         log.NewLog(log.DefaultConfig()),
		Port:        9119,
		TextGateway: nil,
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
	err := http.ListenAndServe(net.JoinHostPort(s.Host, strconv.Itoa(s.Port)), nil)
	if err != nil {
		fmt.Printf("%#v\n", maskAny(err))
	}
}
