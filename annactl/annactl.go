package main

import (
	"github.com/spf13/cobra"

	logcontrol "github.com/xh3b4sd/anna/client/control/log"
	"github.com/xh3b4sd/anna/client/interface/text"
	"github.com/xh3b4sd/anna/file-system/os"
	"github.com/xh3b4sd/anna/spec"
)

func (a *annactl) InitAnnactlCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call InitAnnactlCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "annactl",
		Short: "Interact with Anna's network API. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Interact with Anna's network API. For more information see https://github.com/xh3b4sd/anna.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var err error

			// Log.
			err = a.Log.SetLevels(a.Flags.ControlLogLevels)
			panicOnError(err)
			err = a.Log.SetVerbosity(a.Flags.ControlLogVerbosity)
			panicOnError(err)

			a.Log.Register(a.GetType())

			// File system.
			a.FileSystem = osfilesystem.NewFileSystem(osfilesystem.DefaultConfig())

			// Log control.
			logControlConfig := logcontrol.DefaultControlConfig()
			logControlConfig.URL.Host = a.Flags.HTTPAddr
			a.LogControl, err = logcontrol.NewControl(logControlConfig)
			panicOnError(err)

			// Text interface.
			textInterfaceConfig := text.DefaultClientConfig()
			textInterfaceConfig.GRPCAddr = a.Flags.GRPCAddr
			a.TextInterface, err = text.NewClient(textInterfaceConfig)
			panicOnError(err)
		},
		Run: a.ExecAnnactlCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnactlControlCmd())
	newCmd.AddCommand(a.InitAnnactlInterfaceCmd())
	newCmd.AddCommand(a.InitAnnactlVersionCmd())

	// Define command line flags.
	newCmd.PersistentFlags().StringVar(&a.Flags.GRPCAddr, "grpc-addr", "127.0.0.1:9119", "host:port to bind Anna's gRPC server to")
	newCmd.PersistentFlags().StringVar(&a.Flags.HTTPAddr, "http-addr", "127.0.0.1:9120", "host:port to bind Anna's HTTP server to")
	newCmd.PersistentFlags().StringVarP(&a.Flags.ControlLogLevels, "control-log-levels", "l", "", "set log levels for log control (e.g. E,F)")
	newCmd.PersistentFlags().IntVarP(&a.Flags.ControlLogVerbosity, "control-log-verbosity", "v", 10, "set log verbosity for log control")

	return newCmd
}

func (a *annactl) ExecAnnactlCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call ExecAnnactlCmd")

	cmd.HelpFunc()(cmd, nil)
}
