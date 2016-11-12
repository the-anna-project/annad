package main

import (
	"os"

	"github.com/spf13/cobra"
)

func (a *annad) InitAnnadCmd() *cobra.Command {
	a.Service().Log().Line("func", "InitAnnadCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "annad",
		Short: "Run the anna daemon. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Run the anna daemon. For more information see https://github.com/xh3b4sd/anna.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			a.serviceCollection = a.newServiceCollection()
		},
		Run: a.ExecAnnadCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnadVersionCmd())

	// Define command line flags.
	newCmd.PersistentFlags().StringVar(&a.flags.GRPCAddr, "grpc-addr", "127.0.0.1:9119", "host:port to bind Anna's gRPC server to")
	newCmd.PersistentFlags().StringVar(&a.flags.HTTPAddr, "http-addr", "127.0.0.1:9120", "host:port to bind Anna's HTTP server to")

	newCmd.PersistentFlags().StringVar(&a.flags.StorageType, "storage-type", "redis", "storage type to use for persistency (e.g. memory)")
	newCmd.PersistentFlags().StringVar(&a.flags.RedisConnectionStorageAddr, "redis-connection-storage-addr", "127.0.0.1:6379", "host:port to connect to connection storage")
	newCmd.PersistentFlags().StringVar(&a.flags.RedisConnectionStoragePrefix, "redis-connection-storage-prefix", "anna", "prefix used to prepend to connection storage keys")
	newCmd.PersistentFlags().StringVar(&a.flags.RedisFeatureStorageAddr, "redis-feature-storage-addr", "127.0.0.1:6380", "host:port to connect to feature storage")
	newCmd.PersistentFlags().StringVar(&a.flags.RedisFeatureStoragePrefix, "redis-feature-storage-prefix", "anna", "prefix used to prepend to feature storage keys")
	newCmd.PersistentFlags().StringVar(&a.flags.RedisGeneralStorageAddr, "redis-general-storage-addr", "127.0.0.1:6381", "host:port to connect to general storage")
	newCmd.PersistentFlags().StringVar(&a.flags.RedisGeneralStoragePrefix, "redis-general-storage-prefix", "anna", "prefix used to prepend to general storage keys")

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

	a.Service().Log().Line("msg", "booting metrics endpoint")
	go a.Service().MetricsEndpoint().Boot()

	a.Service().Log().Line("msg", "booting text endpoint")
	go a.Service().TextEndpoint().Boot()

	// Block the main goroutine forever. The process is only supposed to be ended
	// by a call to Shutdown or ForceShutdown.
	select {}
}
