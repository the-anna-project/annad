// Package main implements a command line client for Anna. Cobra CLI is used as
// framework. The commands are simple wrappers around the client package.
package main

import (
	"sync"

	logcontrol "github.com/xh3b4sd/anna/client/control/log"
	"github.com/xh3b4sd/anna/client/interface/text"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/file-system/memory"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectType represents the object type of the command line object.  This is
	// used e.g. to register itself to the logger.
	ObjectType spec.ObjectType = "annactl"
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
	FileSystem    spec.FileSystem
	IDFactory     spec.IDFactory
	Log           spec.Log
	LogControl    spec.LogControl
	TextInterface spec.TextInterfaceClient

	// Settings.
	Flags     Flags
	SessionID string
	Version   string
}

// DefaultConfig provides a default configuration to create a new command line
// object by best effort.
func DefaultConfig() Config {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		panic(err)
	}

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
		FileSystem:    memoryfilesystem.NewFileSystem(memoryfilesystem.DefaultConfig()),
		IDFactory:     newIDFactory,
		Log:           log.NewLog(log.DefaultConfig()),
		LogControl:    newLogControl,
		TextInterface: newTextInterface,

		// Settings.
		Flags:     Flags{},
		SessionID: string(id.MustNew()),
		Version:   version,
	}

	return newConfig
}

// New creates a new configured command line object.
func New(config Config) (spec.Annactl, error) {
	// annactl
	newAnnactl := &annactl{
		Config: config,

		BootOnce:     sync.Once{},
		Closer:       make(chan struct{}, 1),
		ID:           id.MustNew(),
		ShutdownOnce: sync.Once{},
		Type:         spec.ObjectType(ObjectType),
	}

	if newAnnactl.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}

	return newAnnactl, nil
}

type annactl struct {
	Config

	BootOnce     sync.Once
	Closer       chan struct{}
	ID           spec.ObjectID
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (a *annactl) Boot() {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call Boot")

	a.BootOnce.Do(func() {
		go a.listenToSignal()

		a.InitAnnactlCmd().Execute()
	})
}

func (a *annactl) Shutdown() {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call Shutdown")

	a.ShutdownOnce.Do(func() {
		close(a.Closer)
	})
}

func main() {
	newAnnactl, err := New(DefaultConfig())
	panicOnError(err)

	newAnnactl.Boot()
}
