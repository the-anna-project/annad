// Package main implements a command line interface for the anna daemon. Cobra
// CLI is used as framework.
package main

import (
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/network"
	"github.com/xh3b4sd/anna/server"
	"github.com/xh3b4sd/anna/service"
	"github.com/xh3b4sd/anna/service/id"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage"
	storagespec "github.com/xh3b4sd/anna/storage/spec"
)

const (
	// ObjectType represents the object type of the annad object. This is used
	// e.g. to register itself to the logger.
	ObjectType systemspec.ObjectType = "annad"
)

var (
	// version is the project version. It is given via buildflags that inject the
	// commit hash.
	version string
)

// Config represents the configuration used to create a new annad object.
type Config struct {
	// Dependencies.
	Log               systemspec.Log
	Network           systemspec.Network
	Server            systemspec.Server
	ServiceCollection servicespec.Collection
	StorageCollection storagespec.Collection

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
		Log:               log.New(log.DefaultConfig()),
		Network:           network.MustNew(),
		Server:            newServer,
		ServiceCollection: service.MustNewCollection(),
		// TODO remove storage collection
		StorageCollection: storage.MustNewCollection(),

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

	if newAnna.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newAnna.Network == nil {
		return nil, maskAnyf(invalidConfigError, "network must not be empty")
	}
	if newAnna.Server == nil {
		return nil, maskAnyf(invalidConfigError, "server must not be empty")
	}
	if newAnna.ServiceCollection == nil {
		return nil, maskAnyf(invalidConfigError, "service collection must not be empty")
	}
	if newAnna.StorageCollection == nil {
		return nil, maskAnyf(invalidConfigError, "storage collection must not be empty")
	}

	return newAnna, nil
}

type annad struct {
	Config

	BootOnce     sync.Once
	ID           string
	ShutdownOnce sync.Once
	Type         systemspec.ObjectType
}

func (a *annad) Boot() {
	a.Log.WithTags(systemspec.Tags{C: nil, L: "D", O: a, V: 13}, "call Boot")

	a.BootOnce.Do(func() {
		go a.listenToSignal()
		go a.writeStateInfo()

		a.InitAnnadCmd().Execute()
	})
}

func (a *annad) ForceShutdown() {
	a.Log.WithTags(systemspec.Tags{C: nil, L: "D", O: a, V: 13}, "call ForceShutdown")

	a.Log.WithTags(systemspec.Tags{C: nil, L: "I", O: a, V: 10}, "force shutting down annad")
	os.Exit(0)
}

func (a *annad) Service() servicespec.Collection {
	return a.ServiceCollection
}

func (a *annad) Shutdown() {
	a.Log.WithTags(systemspec.Tags{C: nil, L: "D", O: a, V: 13}, "call Shutdown")

	a.ShutdownOnce.Do(func() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			a.Log.WithTags(systemspec.Tags{C: nil, L: "I", O: a, V: 10}, "shutting down service collection")
			a.Service().Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			a.Log.WithTags(systemspec.Tags{C: nil, L: "I", O: a, V: 10}, "shutting down network")
			a.Network.Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			a.Log.WithTags(systemspec.Tags{C: nil, L: "I", O: a, V: 10}, "shutting down server")
			a.Server.Shutdown()
			wg.Done()
		}()

		wg.Wait()

		a.Log.WithTags(systemspec.Tags{C: nil, L: "I", O: a, V: 10}, "shutting down annad")
		os.Exit(0)
	})
}

func (a *annad) Storage() storagespec.Collection {
	return a.StorageCollection
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	newAnna, err := New(DefaultConfig())
	panicOnError(err)

	newAnna.Boot()
}
