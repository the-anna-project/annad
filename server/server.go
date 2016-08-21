// Package server implements a HTTP server to provide Anna's API over network.
package server

import (
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/tylerb/graceful"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/factory/id"
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
	TextInterface   api.TextInterfaceServer

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

	newLogControl, err := logcontrol.NewControl(logcontrol.DefaultControlConfig())
	if err != nil {
		panic(err)
	}

	newTextInterface, err := text.NewInterface(text.DefaultInterfaceConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.

		Instrumentation: newInstrumentation,
		Log:             log.NewLog(log.DefaultConfig()),
		LogControl:      newLogControl,
		TextInterface:   newTextInterface,

		// Settings.

		GRPCAddr: "127.0.0.1:9119",
		HTTPAddr: "127.0.0.1:9120",
	}

	return newConfig
}

// New creates a new configured server object.
func New(config Config) (spec.Server, error) {
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
		ID:           id.MustNew(),
		Mutex:        sync.Mutex{},
		ShutdownOnce: sync.Once{},
		Type:         spec.ObjectType(ObjectTypeServer),
	}

	// Dependencies.

	if newServer.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newServer.LogControl == nil {
		return nil, maskAnyf(invalidConfigError, "log control must not be empty")
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

	newServer.Log.Register(newServer.GetType())

	return newServer, nil
}

type server struct {
	Config

	BootOnce     sync.Once
	Closer       chan struct{}
	GRPCServer   *grpc.Server
	HTTPServer   *graceful.Server
	ID           spec.ObjectID
	Mutex        sync.Mutex
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
		api.RegisterTextInterfaceServer(s.GRPCServer, s.TextInterface)

		// gRPC server.
		fail := make(chan error, 1)
		go func() {
			select {
			case <-s.Closer:
			case err := <-fail:
				s.Log.WithTags(spec.Tags{L: "E", O: s, T: nil, V: 4}, "%#v", maskAny(err))
			}
		}()
		go func() {
			s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 14}, "gRPC server starts to listen on '%s'", s.GRPCAddr)
			listener, err := net.Listen("tcp", s.GRPCAddr)
			if err != nil {
				s.Log.WithTags(spec.Tags{L: "F", O: s, T: nil, V: 1}, "%#v", maskAny(err))
			}
			err = s.GRPCServer.Serve(listener)
			if err != nil {
				fail <- maskAny(err)
			}
		}()

		// HTTP server.
		go func() {
			s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 14}, "HTTP server starts to listen on '%s'", s.HTTPAddr)
			err := s.HTTPServer.ListenAndServe()
			if err != nil {
				s.Log.WithTags(spec.Tags{L: "E", O: s, T: nil, V: 4}, "%#v", maskAny(err))
			}
		}()
	})
}

func (s *server) Shutdown() {
	s.Log.WithTags(spec.Tags{L: "D", O: s, T: nil, V: 13}, "call Shutdown")

	s.ShutdownOnce.Do(func() {
		close(s.Closer)

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			// Stop the gRPC server gracefully and wait some time for open
			// connections to be closed. Then force it to be stopped.
			//
			// TODO we are not stopping the server gracefully right now due to this
			// issue: https://github.com/grpc/grpc-go/issues/848.
			//go s.GRPCServer.GracefulStop()
			//time.Sleep(3 * time.Second)
			//
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
