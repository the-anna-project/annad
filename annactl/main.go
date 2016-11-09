// Package main implements a command line client for Anna. Cobra CLI is used as
// framework. The commands are simple wrappers around the client package.
package main

import (
	"sync"

	"github.com/xh3b4sd/anna/client/interface/text"
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

// Config represents the configuration used to create a new command line
// object.
type Config struct {
	// Dependencies.
	ServiceCollection servicespec.Collection
	TextInterface     systemspec.TextInterfaceClient

	// Settings.
	Flags     Flags
	SessionID string
	Version   string
}

// DefaultConfig provides a default configuration to create a new command line
// object by best effort.
func DefaultConfig() Config {
	newTextInterface, err := text.NewClient(text.DefaultClientConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.
		ServiceCollection: service.MustNewCollection(),
		TextInterface:     newTextInterface,

		// Settings.
		Flags:     Flags{},
		SessionID: string(id.MustNewID()),
		Version:   version,
	}

	return newConfig
}

// New creates a new configured command line object.
func New(config Config) (systemspec.Annactl, error) {
	// annactl
	newAnnactl := &annactl{
		Config: config,

		BootOnce:     sync.Once{},
		Closer:       make(chan struct{}, 1),
		ShutdownOnce: sync.Once{},
	}

	if newAnnactl.ServiceCollection == nil {
		return nil, maskAnyf(invalidConfigError, "service collection must not be empty")
	}

	return newAnnactl, nil
}

type annactl struct {
	Config

	BootOnce     sync.Once
	Closer       chan struct{}
	ShutdownOnce sync.Once
}

func (a *annactl) Boot() {
	a.Service().Log().Line("func", "Boot")

	a.BootOnce.Do(func() {
		go a.listenToSignal()

		a.InitAnnactlCmd().Execute()
	})
}

func (a *annactl) Service() servicespec.Collection {
	return a.ServiceCollection
}

func (a *annactl) Shutdown() {
	a.Service().Log().Line("func", "Shutdown")

	a.ShutdownOnce.Do(func() {
		close(a.Closer)
	})
}

func main() {
	newAnnactl, err := New(DefaultConfig())
	panicOnError(err)

	newAnnactl.Boot()
}
