// Package metric implements a HTTP server to provide Anna's metrics
// over network.
package metric

import (
	"net/http"
	"sync"
	"time"

	"github.com/tylerb/graceful"

	servicespec "github.com/xh3b4sd/anna/service/spec"
)

// New creates a new metric endpoint service.
func New() servicespec.Endpoint {
	return &service{}
}

type service struct {
	// Dependencies.

	serviceCollection servicespec.Collection

	// Settings.

	// address is the host:port representation based on the golang convention for
	// http.ListenAndServe to serve HTTP traffic.
	address      string
	bootOnce     sync.Once
	closer       chan struct{}
	httpServer   *graceful.Server
	metadata     map[string]string
	shutdownOnce sync.Once
}

func (s *service) Boot() {
	s.bootOnce.Do(func() {
		id, err := s.Service().ID().New()
		if err != nil {
			panic(err)
		}
		s.metadata = map[string]string{
			"id":   id,
			"kind": "metric",
			"name": "endpoint",
			"type": "service",
		}

		s.bootOnce = sync.Once{}
		s.closer = make(chan struct{}, 1)
		s.httpServer = &graceful.Server{
			NoSignalHandling: true,
			Server: &http.Server{
				Addr: s.address,
			},
			Timeout: 3 * time.Second,
		}
		s.shutdownOnce = sync.Once{}

		http.Handle(s.Service().Instrumentor().GetHTTPEndpoint(), s.Service().Instrumentor().GetHTTPHandler())

		go func() {
			s.Service().Log().Line("msg", "HTTP server starts to listen on '%s'", s.address)
			err := s.httpServer.ListenAndServe()
			if err != nil {
				s.Service().Log().Line("msg", "%#v", maskAny(err))
			}
		}()
	})
}

func (s *service) Metadata() map[string]string {
	return s.metadata
}

func (s *service) Service() servicespec.Collection {
	return s.serviceCollection
}

func (s *service) SetAddress(address string) {
	s.address = address
}

func (s *service) SetServiceCollection(sc servicespec.Collection) {
	s.serviceCollection = sc
}

func (s *service) Shutdown() {
	s.Service().Log().Line("func", "Shutdown")

	s.shutdownOnce.Do(func() {
		close(s.closer)

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			// Stop the HTTP server gracefully and wait some time for open
			// connections to be closed. Then force it to be stopped.
			s.httpServer.Stop(s.httpServer.Timeout)
			<-s.httpServer.StopChan()
			wg.Done()
		}()

		wg.Wait()
	})
}
