// Package main implements a command line client for Anna. Cobra CLI is used as
// framework. The commands are simple wrappers around the client package.
package main

import (
	"net"
	"sync"

	"github.com/spf13/cobra"

	logcontrol "github.com/xh3b4sd/anna/client/control/log"
	"github.com/xh3b4sd/anna/client/interface/text"
	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/file-system/memory"
	"github.com/xh3b4sd/anna/file-system/os"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	// ObjectTypeAnnactl represents the object type of the command line object.
	// This is used e.g. to register itself to the logger.
	ObjectTypeAnnactl spec.ObjectType = "annactl"
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
	TextInterface spec.TextInterface

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
	newID, err := newIDFactory.WithType(id.Hex128)
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
		FileSystem:    memoryfilesystem.NewFileSystem(memoryfilesystem.DefaultConfig()),
		IDFactory:     newIDFactory,
		Log:           log.NewLog(log.DefaultConfig()),
		LogControl:    newLogControl,
		TextInterface: newTextInterface,

		SessionID: string(newID),
		Version:   version,
	}

	return newConfig
}

// New creates a new configured command line object.
func New(config Config) (spec.Annactl, error) {
	newIDFactory, err := id.NewFactory(id.DefaultFactoryConfig())
	if err != nil {
		return nil, maskAny(err)
	}
	newID, err := newIDFactory.WithType(id.Hex128)
	if err != nil {
		return nil, maskAny(err)
	}

	// annactl
	newAnnactl := &annactl{
		Config: config,

		BootOnce: sync.Once{},
		ID:       newID,
		Type:     spec.ObjectType(ObjectTypeAnnactl),
	}

	if newAnnactl.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}

	// command
	newAnnactl.Cmd = &cobra.Command{
		Use:   "annactl",
		Short: "Interact with Anna's network API. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Interact with Anna's network API. For more information see https://github.com/xh3b4sd/anna.",

		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, nil)
		},

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var err error

			// Log.
			err = newAnnactl.Log.SetLevels(newAnnactl.Flags.ControlLogLevels)
			panicOnError(err)
			err = newAnnactl.Log.SetVerbosity(newAnnactl.Flags.ControlLogVerbosity)
			panicOnError(err)

			newAnnactl.Log.Register(newAnnactl.GetType())

			// File system.
			newFileSystem := osfilesystem.NewFileSystem(osfilesystem.DefaultConfig())

			// host and port.
			host, port, err := net.SplitHostPort(newAnnactl.Flags.Addr)
			panicOnError(err)
			hostport := net.JoinHostPort(host, port)

			// Log control.
			newLogControlConfig := logcontrol.DefaultControlConfig()
			newLogControlConfig.URL.Host = hostport
			newLogControl, err := logcontrol.NewControl(newLogControlConfig)
			panicOnError(err)

			// Text interface.
			newTextInterfaceConfig := text.DefaultInterfaceConfig()
			newTextInterfaceConfig.URL.Host = hostport
			newTextInterface, err := text.NewInterface(newTextInterfaceConfig)
			panicOnError(err)

			// Annactl.
			newAnnactl.FileSystem = newFileSystem
			newAnnactl.LogControl = newLogControl
			newAnnactl.TextInterface = newTextInterface
		},
	}

	// Flags.
	newAnnactl.Cmd.PersistentFlags().StringVar(&newAnnactl.Flags.Addr, "addr", "127.0.0.1:9119", "host:port to connect to Anna's server")

	newAnnactl.Cmd.PersistentFlags().StringVarP(&newAnnactl.Flags.ControlLogLevels, "control-log-levels", "l", "", "set log levels for log control (e.g. E,F)")
	newAnnactl.Cmd.PersistentFlags().IntVarP(&newAnnactl.Flags.ControlLogVerbosity, "control-log-verbosity", "v", 10, "set log verbosity for log control")

	return newAnnactl, nil
}

type annactl struct {
	Config

	Cmd      *cobra.Command
	BootOnce sync.Once
	ID       spec.ObjectID
	Type     spec.ObjectType
}

func (a *annactl) Boot() {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call Boot")

	a.BootOnce.Do(func() {
		// init
		a.Cmd.AddCommand(a.InitControlCmd())
		a.Cmd.AddCommand(a.InitInterfaceCmd())
		a.Cmd.AddCommand(a.InitVersionCmd())

		// execute
		a.Cmd.Execute()
	})
}

func main() {
	newAnnactl, err := New(DefaultConfig())
	panicOnError(err)

	newAnnactl.Boot()
}
