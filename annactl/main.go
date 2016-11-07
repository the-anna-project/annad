// Package main implements a command line client for Anna. Cobra CLI is used as
// framework. The commands are simple wrappers around the client package.
package main

import (
	"sync"

	logcontrol "github.com/xh3b4sd/anna/client/control/log"
	"github.com/xh3b4sd/anna/client/interface/text"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/service"
	"github.com/xh3b4sd/anna/service/id"
	servicespec "github.com/xh3b4sd/anna/service/spec"
	systemspec "github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectType represents the object type of the command line object.  This is
	// used e.g. to register itself to the logger.
	ObjectType systemspec.ObjectType = "annactl"
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
	Log               systemspec.Log
	LogControl        systemspec.LogControl
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
	newLogControl, err := logcontrol.NewControl(logcontrol.DefaultControlConfig())
	if err != nil {
		panic(err)
	}

	newTextInterface, err := text.NewClient(text.DefaultClientConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.
		Log:               log.New(log.DefaultConfig()),
		LogControl:        newLogControl,
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
		ID:           id.MustNewID(),
		ShutdownOnce: sync.Once{},
		Type:         systemspec.ObjectType(ObjectType),
	}

	if newAnnactl.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
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
	ID           string
	ShutdownOnce sync.Once
	Type         systemspec.ObjectType
}

func (a *annactl) Boot() {
	a.Log.WithTags(systemspec.Tags{C: nil, L: "D", O: a, V: 13}, "call Boot")

	a.BootOnce.Do(func() {
		go a.listenToSignal()

		a.InitAnnactlCmd().Execute()
	})
}

func (a *annactl) Service() servicespec.Collection {
	return a.ServiceCollection
}

func (a *annactl) Shutdown() {
	a.Log.WithTags(systemspec.Tags{C: nil, L: "D", O: a, V: 13}, "call Shutdown")

	a.ShutdownOnce.Do(func() {
		close(a.Closer)
	})
}

func main() {
	newAnnactl, err := New(DefaultConfig())
	panicOnError(err)

	newAnnactl.Boot()
}
