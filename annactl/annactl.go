// Package annactl implements a command line client for Anna. Cobra CLI is used
// as framework. The commands are simple wrappers around the client package.
package main

import (
	"net"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/client/control/log"
	"github.com/xh3b4sd/anna/client/interface/text"
	"github.com/xh3b4sd/anna/id"
	logpkg "github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	objectTypeAnnactl spec.ObjectType = "annactl"
)

var (
	a spec.Object

	globalFlags struct {
		Addr string
	}

	log           spec.Log
	logControl    spec.LogControl
	textInterface spec.TextInterface

	mainCmd = &cobra.Command{
		Use:   "annactl",
		Short: "Interact with Anna's network API. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Interact with Anna's network API. For more information see https://github.com/xh3b4sd/anna.",
		Run:   mainRun,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// annactl
			a = newAnnactl(defaultAnnactlConfig())

			// log
			log = logpkg.NewLog(logpkg.DefaultConfig())
			log.Register(a.GetType())

			// host and port
			host, port, err := net.SplitHostPort(globalFlags.Addr)
			panicOnError(err)
			hostport := net.JoinHostPort(host, port)

			// log control
			newLogControlConfig := logcontrol.DefaultConfig()
			newLogControlConfig.URL.Host = hostport
			logControl = logcontrol.NewLogControl(newLogControlConfig)

			// text interface
			newTextInterfaceConfig := textinterface.DefaultConfig()
			newTextInterfaceConfig.URL.Host = hostport
			textInterface = textinterface.NewTextInterface(newTextInterfaceConfig)
		},
	}

	// Version is the project version. It is given via buildflags that inject the
	// commit hash.
	version string
)

func init() {
	mainCmd.PersistentFlags().StringVar(&globalFlags.Addr, "addr", "127.0.0.1:9119", "host:port to connect to Anna's server")
}

type annactlConfig struct{}

func defaultAnnactlConfig() annactlConfig {
	newConfig := annactlConfig{}

	return newConfig
}

func newAnnactl(config annactlConfig) spec.Object {
	newAnnactl := &annactl{
		annactlConfig: config,
		ID:            id.NewObjectID(id.Hex128),
		Type:          spec.ObjectType(objectTypeAnnactl),
	}

	return newAnnactl
}

type annactl struct {
	annactlConfig

	ID   spec.ObjectID
	Type spec.ObjectType
}

func mainRun(cmd *cobra.Command, args []string) {
	cmd.Help()
}

func main() {
	controlLogResetCmd.AddCommand(controlLogResetLevelsCmd)
	controlLogResetCmd.AddCommand(controlLogResetObjectsCmd)
	controlLogResetCmd.AddCommand(controlLogResetVerbosityCmd)
	controlLogSetCmd.AddCommand(controlLogSetLevelsCmd)
	controlLogSetCmd.AddCommand(controlLogSetObjectsCmd)
	controlLogSetCmd.AddCommand(controlLogSetVerbosityCmd)
	controlLogCmd.AddCommand(controlLogResetCmd)
	controlLogCmd.AddCommand(controlLogSetCmd)
	controlCmd.AddCommand(controlLogCmd)
	mainCmd.AddCommand(controlCmd)

	interfaceTextReadCmd.AddCommand(interfaceTextReadPlainCmd)
	interfaceTextCmd.AddCommand(interfaceTextReadCmd)
	interfaceCmd.AddCommand(interfaceTextCmd)
	mainCmd.AddCommand(interfaceCmd)

	mainCmd.AddCommand(versionCmd)

	mainCmd.Execute()
}
