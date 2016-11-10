package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/server"
)

func (a *annad) InitAnnadCmd() *cobra.Command {
	a.Service().Log().Line("func", "InitAnnadCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "annad",
		Short: "Run the anna daemon. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Run the anna daemon. For more information see https://github.com/xh3b4sd/anna.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// service collection.
			a.serviceCollection = newServiceCollection()

			// TODO move storage to service collection
			// storage collection.
			_, err := newStorageCollection(a.flags)
			panicOnError(err)

			// text interface.
			textInterface, err := newTextInterface(a.serviceCollection)
			panicOnError(err)

			// server.
			serverConfig := server.DefaultConfig()
			serverConfig.GRPCAddr = a.flags.GRPCAddr
			serverConfig.HTTPAddr = a.flags.HTTPAddr
			serverConfig.Instrumentation, err = newPrometheusInstrumentation([]string{"Server"})
			panicOnError(err)
			serverConfig.ServiceCollection = a.serviceCollection
			serverConfig.TextInterface = textInterface
			a.server, err = server.New(serverConfig)
			panicOnError(err)
		},
		Run: a.ExecAnnadCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnadVersionCmd())

	// Define command line flags.
	newCmd.PersistentFlags().StringVar(&a.flags.GRPCAddr, "grpc-addr", "127.0.0.1:9119", "host:port to bind Anna's gRPC server to")
	newCmd.PersistentFlags().StringVar(&a.flags.HTTPAddr, "http-addr", "127.0.0.1:9120", "host:port to bind Anna's HTTP server to")

	newCmd.PersistentFlags().StringVar(&a.flags.ControlLogLevels, "control-log-levels", "", "set log levels for log control (e.g. E)")
	newCmd.PersistentFlags().StringVar(&a.flags.ControlLogObejcts, "control-log-objects", "", "set log objects for log control (e.g. network)")
	newCmd.PersistentFlags().IntVar(&a.flags.ControlLogVerbosity, "control-log-verbosity", 10, "set log verbosity for log control")

	newCmd.PersistentFlags().StringVar(&a.flags.Storage, "storage", "redis", "storage type to use for persistency (e.g. memory)")
	newCmd.PersistentFlags().StringVar(&a.flags.RedisFeatureStorageAddr, "redis-feature-storage-addr", "127.0.0.1:6380", "host:port to connect to feature storage")
	newCmd.PersistentFlags().StringVar(&a.flags.RedisGeneralStorageAddr, "redis-general-storage-addr", "127.0.0.1:6381", "host:port to connect to general storage")
	newCmd.PersistentFlags().StringVar(&a.flags.RedisStoragePrefix, "redis-storage-prefix", "anna", "prefix used to prepend to storage keys")

	return newCmd
}

func (a *annad) ExecAnnadCmd(cmd *cobra.Command, args []string) {
	a.Service().Log().Line("func", "ExecAnnadCmd")

	if len(args) > 0 {
		cmd.HelpFunc()(cmd, nil)
		os.Exit(1)
	}

	a.Service().Log().Line("msg", "booting annad")

	a.Service().Log().Line("msg", "booting network")
	go a.Service().Network().Boot()

	a.Service().Log().Line("msg", "booting server")
	go a.server.Boot()

	// Block the main goroutine forever. The process is only supposed to be ended
	// by a call to Shutdown or ForceShutdown.
	select {}
}
