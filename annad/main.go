// Package main implements a command line interface for the anna daemon. Cobra
// CLI is used as framework.
package main

import (
	"math/rand"
	"os"
	"sync"
	"time"

	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
)

var (
	// version is the project version. It is given via buildflags that inject the
	// commit hash.
	version string
)

// New creates a new annad service.
func New() systemspec.Annad {
	return &annad{}
}

type annad struct {
	// Dependencies.

	server            systemspec.Server
	serviceCollection servicespec.Collection

	// Settings.

	bootOnce     sync.Once
	flags        Flags
	metadata     map[string]string
	shutdownOnce sync.Once
	version      string
}

func (a *annad) Boot() {
	a.Service().Log().Line("func", "Boot")

	a.bootOnce.Do(func() {
		go a.listenToSignal()

		a.InitAnnadCmd().Execute()
	})
}

func (a *annad) Configure() error {
	// Settings.

	id, err := a.Service().ID().New()
	if err != nil {
		return maskAny(err)
	}
	a.metadata = map[string]string{
		"id":   id,
		"name": "annad",
		"type": "service",
	}

	a.bootOnce = sync.Once{}
	a.flags = Flags{}
	a.shutdownOnce = sync.Once{}
	a.version = version

	return nil
}

func (a *annad) ForceShutdown() {
	a.Service().Log().Line("func", "ForceShutdown")

	a.Service().Log().Line("msg", "force shutting down annad")
	os.Exit(0)
}

func (a *annad) Metadata() map[string]string {
	return a.metadata
}

func (a *annad) Service() servicespec.Collection {
	return a.serviceCollection
}

func (a *annad) SetServer(s systemspec.Server) {
	a.server = s
}

func (a *annad) SetServiceCollection(sc servicespec.Collection) {
	a.serviceCollection = sc
}

func (a *annad) Shutdown() {
	a.Service().Log().Line("func", "Shutdown")

	a.shutdownOnce.Do(func() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			a.Service().Log().Line("msg", "shutting down service collection")
			a.Service().Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			a.Service().Log().Line("msg", "shutting down server")
			a.server.Shutdown()
			wg.Done()
		}()

		wg.Wait()

		a.Service().Log().Line("msg", "shutting down annad")
		os.Exit(0)
	})
}

func (a *annad) Validate() error {
	// Dependencies.
	if a.server == nil {
		return maskAnyf(invalidConfigError, "server must not be empty")
	}
	if a.serviceCollection == nil {
		return maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	return nil
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	newAnnad := New()

	newAnnad.Boot()
}
