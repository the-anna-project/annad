// Package annactl implements a command line client for Anna. Cobra CLI is used
// as framework. The commands are simple wrappers around the client package.
package main

import (
	"net"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/client/control/log"
	"github.com/xh3b4sd/anna/client/interface/text"
	"github.com/xh3b4sd/anna/file-system/memory"
	"github.com/xh3b4sd/anna/file-system/os"
	"github.com/xh3b4sd/anna/id"
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
	Log           spec.Log
	LogControl    spec.LogControl
	TextInterface spec.TextInterface

	// Settings.
	Cmd       *cobra.Command
	Flags     Flags
	SessionID string
	Version   string
}

// DefaultConfig provides a default configuration to create a new command line
// object by best effort.
func DefaultConfig() Config {
	newConfig := Config{
		FileSystem:    memoryfilesystem.NewFileSystem(),
		Log:           log.NewLog(log.DefaultConfig()),
		LogControl:    logcontrol.NewLogControl(logcontrol.DefaultConfig()),
		TextInterface: textinterface.NewTextInterface(textinterface.DefaultConfig()),

		SessionID: string(id.NewObjectID(id.Hex128)),
		Version:   version,
	}

	return newConfig
}

// NewAnnactl creates a new configured command line object.
func NewAnnactl(config Config) spec.Annactl {
	// annactl
	newAnnactl := &annactl{
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Type:   spec.ObjectType(ObjectTypeAnnactl),
	}

	// command
	newAnnactl.Cmd = &cobra.Command{
		Use:   "annactl",
		Short: "Interact with Anna's network API. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Interact with Anna's network API. For more information see https://github.com/xh3b4sd/anna.",

		Run: func(cmd *cobra.Command, args []string) {
			newAnnactl.Log.WithTags(spec.Tags{L: "D", O: newAnnactl, T: nil, V: 13}, "call ExecCmd") // this is the first stage we can log

			cmd.Help()
		},

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var err error

			// log
			newLog := log.NewLog(log.DefaultConfig())
			err = newLog.SetLevels(newAnnactl.Flags.ControlLogLevels)
			panicOnError(err)
			err = newLog.SetVerbosity(newAnnactl.Flags.ControlLogVerbosity)
			panicOnError(err)

			newAnnactl.Log.WithTags(spec.Tags{L: "D", O: newAnnactl, T: nil, V: 13}, "call InitCmd") // this is the first stage we can log

			// file system
			newFileSystem := osfilesystem.NewFileSystem()

			// host and port
			host, port, err := net.SplitHostPort(newAnnactl.Flags.Addr)
			panicOnError(err)
			hostport := net.JoinHostPort(host, port)

			// log control
			newLogControlConfig := logcontrol.DefaultConfig()
			newLogControlConfig.URL.Host = hostport
			newLogControl := logcontrol.NewLogControl(newLogControlConfig)

			// text interface
			newTextInterfaceConfig := textinterface.DefaultConfig()
			newTextInterfaceConfig.URL.Host = hostport
			newTextInterface := textinterface.NewTextInterface(newTextInterfaceConfig)

			// annactl
			newAnnactl.FileSystem = newFileSystem
			newAnnactl.Log = newLog
			newAnnactl.LogControl = newLogControl
			newAnnactl.TextInterface = newTextInterface

			newAnnactl.Log.Register(newAnnactl.GetType())
		},
	}

	// flags
	newAnnactl.Cmd.PersistentFlags().StringVar(&newAnnactl.Flags.Addr, "addr", "127.0.0.1:9119", "host:port to connect to Anna's server")

	newAnnactl.Cmd.PersistentFlags().StringVarP(&newAnnactl.Flags.ControlLogLevels, "control-log-levels", "l", "", "set log levels for log control (e.g. E,F)")
	newAnnactl.Cmd.PersistentFlags().IntVarP(&newAnnactl.Flags.ControlLogVerbosity, "control-log-verbosity", "v", 10, "set log verbosity for log control")

	return newAnnactl
}

func (a *annactl) Boot() {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call Boot")

	// init
	a.Cmd.AddCommand(a.InitControlCmd())
	a.Cmd.AddCommand(a.InitInterfaceCmd())
	a.Cmd.AddCommand(a.InitVersionCmd())

	// execute
	a.Cmd.Execute()
}

type annactl struct {
	Config

	ID   spec.ObjectID
	Type spec.ObjectType
}

func main() {
	newAnnactl := NewAnnactl(DefaultConfig())
	newAnnactl.Boot()
}
