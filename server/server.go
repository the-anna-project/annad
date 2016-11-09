// Package server implements a HTTP server to provide Anna's API over network.
package server

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/tylerb/graceful"
	"google.golang.org/grpc"

	"github.com/xh3b4sd/anna/instrumentation/memory"
	"github.com/xh3b4sd/anna/server/interface/text"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
)

// Config represents the configuration used to create a new server object.
type Config struct {
	// Dependencies.

	Instrumentation   systemspec.Instrumentation
	ServiceCollection servicespec.Collection
	TextInterface     text.TextInterfaceServer

	// Settings.

	// GRPCAddr is the host:port representation based on the golang convention
	// for net.Listen to serve gRPC traffic.
	GRPCAddr string

	// HTTPAddr is the host:port representation based on the golang convention
	// for http.ListenAndServe to serve HTTP traffic.
	HTTPAddr string
}

// DefaultConfig provides a default configuration to create a new server object
// by best effort.
func DefaultConfig() Config {
	newInstrumentation, err := memory.NewInstrumentation(memory.DefaultInstrumentationConfig())
	if err != nil {
		panic(err)
	}

	newTextInterface, err := text.NewServer(text.DefaultServerConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.

		Instrumentation:   newInstrumentation,
		ServiceCollection: nil,
		TextInterface:     newTextInterface,

		// Settings.

		GRPCAddr: "127.0.0.1:9119",
		HTTPAddr: "127.0.0.1:9120",
	}

	return newConfig
}

// New creates a new configured server object.
func New(config Config) (systemspec.Server, error) {
	newServer := &server{
		Config: config,

		BootOnce:   sync.Once{},
		Closer:     make(chan struct{}, 1),
		GRPCServer: grpc.NewServer(),
		HTTPServer: &graceful.Server{
			NoSignalHandling: true,
			Server: &http.Server{
				Addr: config.HTTPAddr,
			},
			Timeout: 3 * time.Second,
		},
		ShutdownOnce: sync.Once{},
	}

	// Dependencies.
	if newServer.ServiceCollection == nil {
		return nil, maskAnyf(invalidConfigError, "service collection must not be empty")
	}
	if newServer.TextInterface == nil {
		return nil, maskAnyf(invalidConfigError, "text interface must not be empty")
	}

	// Settings.
	if newServer.GRPCAddr == "" {
		return nil, maskAnyf(invalidConfigError, "gRPC address must not be empty")
	}
	if newServer.HTTPAddr == "" {
		return nil, maskAnyf(invalidConfigError, "HTTP address must not be empty")
	}

	id, err := newServer.Service().ID().New()
	if err != nil {
		return nil, maskAny(err)
	}
	newServer.Metadata["id"] = id
	newServer.Metadata["name"] = "server"
	newServer.Metadata["type"] = "service"

	return newServer, nil
}

type server struct {
	Config

	BootOnce     sync.Once
	Closer       chan struct{}
	GRPCServer   *grpc.Server
	HTTPServer   *graceful.Server
	Metadata     map[string]string
	ShutdownOnce sync.Once
}

func (s *server) Boot() {
	s.Service().Log().Line("func", "Boot")

	s.BootOnce.Do(func() {
		// Instrumentation.
		http.Handle(s.Instrumentation.GetHTTPEndpoint(), s.Instrumentation.GetHTTPHandler())

		// Text interface.
		text.RegisterTextInterfaceServer(s.GRPCServer, s.TextInterface)

		// Create the gRPC server. The Serve method below is returning listener
		// errors, if any. In case net.Listener.Accept is called and waits for
		// connections while the listener was closed, a net.OpError will be thrown.
		// For this case we only log errors from the fail channel in case the
		// server's Closer was not closed yet.
		fail := make(chan error, 1)
		go func() {
			select {
			case <-s.Closer:
			case err := <-fail:
				s.Service().Log().Line("msg", "%#v", maskAny(err))
			}
		}()
		go func() {
			s.Service().Log().Line("msg", "gRPC server starts to listen on '%s'", s.GRPCAddr)
			listener, err := net.Listen("tcp", s.GRPCAddr)
			if err != nil {
				s.Service().Log().Line("msg", "%#v", maskAny(err))
			}
			err = s.GRPCServer.Serve(listener)
			if err != nil {
				fail <- maskAny(err)
			}
		}()

		// HTTP server.
		go func() {
			s.Service().Log().Line("msg", "HTTP server starts to listen on '%s'", s.HTTPAddr)
			err := s.HTTPServer.ListenAndServe()
			if err != nil {
				s.Service().Log().Line("msg", "%#v", maskAny(err))
			}
		}()
	})
}

func (s *server) Shutdown() {
	s.Service().Log().Line("func", "Shutdown")

	s.ShutdownOnce.Do(func() {
		close(s.Closer)

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			// Stop the gRPC server gracefully and wait some time for open
			// connections to be closed. Then force it to be stopped.
			go s.GRPCServer.GracefulStop()
			time.Sleep(3 * time.Second)
			s.GRPCServer.Stop()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			// Stop the HTTP server gracefully and wait some time for open
			// connections to be closed. Then force it to be stopped.
			s.HTTPServer.Stop(s.HTTPServer.Timeout)
			<-s.HTTPServer.StopChan()
			wg.Done()
		}()

		wg.Wait()
	})
}
