// Package main implements a command line interface for the anna daemon. Cobra
// CLI is used as framework.
package main

import (
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/xh3b4sd/anna/network"
	"github.com/xh3b4sd/anna/server"
	"github.com/xh3b4sd/anna/service"
	"github.com/xh3b4sd/anna/service/id"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
)

var (
	// version is the project version. It is given via buildflags that inject the
	// commit hash.
	version string
)

// Config represents the configuration used to create a new annad object.
type Config struct {
	// Dependencies.
	Network           systemspec.Network
	Server            systemspec.Server
	ServiceCollection servicespec.Collection

	// Settings.
	Flags   Flags
	Version string
}

// DefaultConfig provides a default configuration to create a new annad object
// by best effort.
func DefaultConfig() Config {
	newServer, err := server.New(server.DefaultConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.
		Network:           network.MustNew(),
		Server:            newServer,
		ServiceCollection: service.MustNewCollection(),

		// Settings.
		Flags:   Flags{},
		Version: version,
	}

	return newConfig
}

// New creates a new configured annad object.
func New(config Config) (systemspec.Annad, error) {
	newAnna := &annad{
		Config: config,

		BootOnce:     sync.Once{},
		ID:           id.MustNewID(),
		ShutdownOnce: sync.Once{},
		Type:         ObjectType,
	}

	// Dependencies.
	if newAnna.Network == nil {
		return nil, maskAnyf(invalidConfigError, "network must not be empty")
	}
	if newAnna.Server == nil {
		return nil, maskAnyf(invalidConfigError, "server must not be empty")
	}
	if newAnna.ServiceCollection == nil {
		return nil, maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	return newAnna, nil
}

type annad struct {
	Config

	BootOnce     sync.Once
	ID           string
	ShutdownOnce sync.Once
	Type         string
}

func (a *annad) Boot() {
	a.Service().Log().Line("func", "Boot")

	a.BootOnce.Do(func() {
		go a.listenToSignal()

		a.InitAnnadCmd().Execute()
	})
}

func (a *annad) ForceShutdown() {
	a.Service().Log().Line("func", "ForceShutdown")

	a.Service().Log().Line("msg", "force shutting down annad")
	os.Exit(0)
}

func (a *annad) Service() servicespec.Collection {
	return a.ServiceCollection
}

func (a *annad) Shutdown() {
	a.Service().Log().Line("func", "Shutdown")

	a.ShutdownOnce.Do(func() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			a.Service().Log().Line("msg", "shutting down service collection")
			a.Service().Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			a.Service().Log().Line("msg", "shutting down network")
			a.Network.Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			a.Service().Log().Line("msg", "shutting down server")
			a.Server.Shutdown()
			wg.Done()
		}()

		wg.Wait()

		a.Service().Log().Line("msg", "shutting down annad")
		os.Exit(0)
	})
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	newAnna, err := New(DefaultConfig())
	panicOnError(err)

	newAnna.Boot()
}
