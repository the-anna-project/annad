// Package main implements a command line for Anna. Cobra CLI is used as
// framework.
package main

import (
	"os"
	"sync"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/network"
	"github.com/xh3b4sd/anna/server"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectType represents the object type of the anna object. This is used
	// e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "anna"
)

var (
	// version is the project version. It is given via buildflags that inject the
	// commit hash.
	version string
)

// Config represents the configuration used to create a new anna object.
type Config struct {
	// Dependencies.
	FeatureStorage spec.Storage
	GeneralStorage spec.Storage
	Log            spec.Log
	Network        spec.Network
	Server         spec.Server

	// Settings.
	Flags   Flags
	Version string
}

// DefaultConfig provides a default configuration to create a new anna object
// by best effort.
func DefaultConfig() Config {
	newNetwork, err := network.New(network.DefaultConfig())
	if err != nil {
		panic(err)
	}

	newServer, err := server.New(server.DefaultConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.
		FeatureStorage: memory.MustNew(),
		GeneralStorage: memory.MustNew(),
		Log:            log.New(log.DefaultConfig()),
		Network:        newNetwork,
		Server:         newServer,

		// Settings.
		Flags:   Flags{},
		Version: version,
	}

	return newConfig
}

// New creates a new configured anna object.
func New(config Config) (spec.Anna, error) {
	newAnna := &anna{
		Config: config,

		BootOnce:     sync.Once{},
		ID:           id.MustNew(),
		ShutdownOnce: sync.Once{},
		Type:         ObjectType,
	}

	if newAnna.FeatureStorage == nil {
		return nil, maskAnyf(invalidConfigError, "feature storage must not be empty")
	}
	if newAnna.GeneralStorage == nil {
		return nil, maskAnyf(invalidConfigError, "generale storage must not be empty")
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

	return newAnna, nil
}

type anna struct {
	Config

	BootOnce     sync.Once
	ID           spec.ObjectID
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (a *anna) Boot() {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call Boot")

	a.BootOnce.Do(func() {
		go a.listenToSignal()
		go a.writeStateInfo()

		a.InitAnnaCmd().Execute()
	})
}

func (a *anna) ForceShutdown() {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call ForceShutdown")

	a.Log.WithTags(spec.Tags{C: nil, L: "I", O: a, V: 10}, "force shutting down Anna")
	os.Exit(0)
}

func (a *anna) Shutdown() {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call Shutdown")

	a.ShutdownOnce.Do(func() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			a.Log.WithTags(spec.Tags{C: nil, L: "I", O: a, V: 10}, "shutting down network")
			a.Network.Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			a.Log.WithTags(spec.Tags{C: nil, L: "I", O: a, V: 10}, "shutting down server")
			a.Server.Shutdown()
			wg.Done()
		}()

		wg.Wait()

		a.Log.WithTags(spec.Tags{C: nil, L: "I", O: a, V: 10}, "shutting down Anna")
		os.Exit(0)
	})
}

func main() {
	newAnna, err := New(DefaultConfig())
	panicOnError(err)

	newAnna.Boot()
}
