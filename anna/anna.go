package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/network"
	"github.com/xh3b4sd/anna/server"
	"github.com/xh3b4sd/anna/spec"
)

func (a *anna) InitAnnaCmd() *cobra.Command {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call InitAnnaCmd")

	// Create new command.
	newCmd := &cobra.Command{
		Use:   "anna",
		Short: "Anna, Artificial Neural Network Aspiration, aims to be self-learning and self-improving software. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Anna, Artificial Neural Network Aspiration, aims to be self-learning and self-improving software. For more information see https://github.com/xh3b4sd/anna.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var err error

			// log.
			err = a.Log.SetLevels(a.Flags.ControlLogLevels)
			panicOnError(err)
			err = a.Log.SetObjects(a.Flags.ControlLogObejcts)
			panicOnError(err)
			err = a.Log.SetVerbosity(a.Flags.ControlLogVerbosity)
			panicOnError(err)

			a.Log.Register(a.GetType())

			// text input/output channel.
			newTextInput := make(chan spec.TextRequest, 1000)
			newTextOutput := make(chan spec.TextResponse, 1000)

			// factory collection.
			factoryCollection, err := newFactoryCollection()
			panicOnError(err)

			// gateway collection.
			gatewayCollection, err := newGatewayCollection(newTextOutput)
			panicOnError(err)

			// storage collection.
			a.StorageCollection, err = newStorageCollection(a.Log, a.Closer, a.Flags)
			panicOnError(err)

			// activator.
			activator, err := newActivator(a.Log, factoryCollection, a.StorageCollection)
			panicOnError(err)

			// forwarder.
			forwarder, err := newForwarder(a.Log, factoryCollection, a.StorageCollection)
			panicOnError(err)

			// network.
			networkConfig := network.DefaultConfig()
			networkConfig.Activator = activator
			networkConfig.FactoryCollection = factoryCollection
			networkConfig.Forwarder = forwarder
			networkConfig.GatewayCollection = gatewayCollection
			networkConfig.Log = a.Log
			networkConfig.StorageCollection = a.StorageCollection
			networkConfig.TextInput = newTextInput
			networkConfig.TextOutput = newTextOutput
			a.Network, err = network.New(networkConfig)
			panicOnError(err)

			// log control.
			logControl, err := newLogControl(a.Log)
			panicOnError(err)

			// text interface.
			textInterface, err := newTextInterface(a.Log, newTextInput, newTextOutput)
			panicOnError(err)

			// server.
			serverConfig := server.DefaultConfig()
			serverConfig.GRPCAddr = a.Flags.GRPCAddr
			serverConfig.HTTPAddr = a.Flags.HTTPAddr
			serverConfig.Instrumentation, err = newPrometheusInstrumentation([]string{"Server"})
			panicOnError(err)
			serverConfig.Log = a.Log
			serverConfig.LogControl = logControl
			serverConfig.TextInterface = textInterface
			a.Server, err = server.New(serverConfig)
			panicOnError(err)
		},
		Run: a.ExecAnnaCmd,
	}

	// Add sub commands.
	newCmd.AddCommand(a.InitAnnaVersionCmd())

	// Define command line flags.
	newCmd.PersistentFlags().StringVar(&a.Flags.GRPCAddr, "grpc-addr", "127.0.0.1:9119", "host:port to bind Anna's gRPC server to")
	newCmd.PersistentFlags().StringVar(&a.Flags.HTTPAddr, "http-addr", "127.0.0.1:9120", "host:port to bind Anna's HTTP server to")

	newCmd.PersistentFlags().StringVar(&a.Flags.ControlLogLevels, "control-log-levels", "", "set log levels for log control (e.g. E)")
	newCmd.PersistentFlags().StringVar(&a.Flags.ControlLogObejcts, "control-log-objects", "", "set log objects for log control (e.g. network)")
	newCmd.PersistentFlags().IntVar(&a.Flags.ControlLogVerbosity, "control-log-verbosity", 10, "set log verbosity for log control")

	newCmd.PersistentFlags().StringVar(&a.Flags.Storage, "storage", "redis", "storage type to use for persistency (e.g. memory)")
	newCmd.PersistentFlags().StringVar(&a.Flags.RedisFeatureStorageAddr, "redis-feature-storage-addr", "127.0.0.1:6380", "host:port to connect to feature storage")
	newCmd.PersistentFlags().StringVar(&a.Flags.RedisGeneralStorageAddr, "redis-general-storage-addr", "127.0.0.1:6381", "host:port to connect to general storage")
	newCmd.PersistentFlags().StringVar(&a.Flags.RedisStoragePrefix, "redis-storage-prefix", "anna", "prefix used to prepend to storage keys")

	return newCmd
}

func (a *anna) ExecAnnaCmd(cmd *cobra.Command, args []string) {
	a.Log.WithTags(spec.Tags{C: nil, L: "D", O: a, V: 13}, "call ExecAnnaCmd")

	if len(args) > 0 {
		cmd.HelpFunc()(cmd, nil)
		os.Exit(1)
	}

	a.Log.WithTags(spec.Tags{C: nil, L: "I", O: a, V: 10}, "booting Anna")

	a.Log.WithTags(spec.Tags{C: nil, L: "I", O: a, V: 10}, "booting network")
	go a.Network.Boot()

	a.Log.WithTags(spec.Tags{C: nil, L: "I", O: a, V: 10}, "booting server")
	go a.Server.Boot()

	// Block the main goroutine forever. The process is only supposed to be ended
	// by a call to Shutdown or ForceShutdown.
	select {}
}
