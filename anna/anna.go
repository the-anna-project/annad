// Package main implements a command line for Anna. Cobra CLI is used as
// framework.
package main

import (
	"os"
	"sync"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/factory/id"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/network"
	"github.com/xh3b4sd/anna/scheduler"
	"github.com/xh3b4sd/anna/server"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/storage/memory"
)

const (
	// ObjectTypeAnna represents the object type of the anna object. This is used
	// e.g. to register itself to the logger.
	ObjectTypeAnna spec.ObjectType = "anna"
)

var (
	// version is the project version. It is given via buildflags that inject the
	// commit hash.
	version string
)

// Config represents the configuration used to create a new anna object.
type Config struct {
	// Dependencies.
	Network spec.Network
	Log     spec.Log
	Server  spec.Server
	Storage spec.Storage

	// Settings.
	Flags   Flags
	Version string
}

// DefaultConfig provides a default configuration to create a new anna object
// by best effort.
func DefaultConfig() Config {
	newNetwork, err := network.New(network.DefaultConfig())
	if err != nil {
		panic(err)
	}

	newServer, err := server.New(server.DefaultConfig())
	if err != nil {
		panic(err)
	}

	newStorage, err := memory.NewStorage(memory.DefaultStorageConfig())
	if err != nil {
		panic(err)
	}

	newConfig := Config{
		// Dependencies.
		Network: newNetwork,
		Log:     log.NewLog(log.DefaultConfig()),
		Server:  newServer,
		Storage: newStorage,

		// Settings.
		Version: version,
	}

	return newConfig
}

// New creates a new configured anna object.
func New(config Config) (spec.Anna, error) {
	newAnna := &anna{
		Config: config,

		BootOnce:     sync.Once{},
		ID:           id.MustNew(),
		ShutdownOnce: sync.Once{},
		Type:         spec.ObjectType(ObjectTypeAnna),
	}

	if newAnna.Network == nil {
		return nil, maskAnyf(invalidConfigError, "network must not be empty")
	}
	if newAnna.Log == nil {
		return nil, maskAnyf(invalidConfigError, "logger must not be empty")
	}
	if newAnna.Server == nil {
		return nil, maskAnyf(invalidConfigError, "server must not be empty")
	}
	if newAnna.Storage == nil {
		return nil, maskAnyf(invalidConfigError, "storage must not be empty")
	}

	newAnna.Cmd = &cobra.Command{
		Use:   "anna",
		Short: "Anna, Artificial Neural Network Aspiration, aims to be self-learning and self-improving software. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Anna, Artificial Neural Network Aspiration, aims to be self-learning and self-improving software. For more information see https://github.com/xh3b4sd/anna.",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				cmd.Help()
				os.Exit(1)
			}

			// Boot.
			newAnna.Log.WithTags(spec.Tags{L: "I", O: newAnna, T: nil, V: 10}, "booting Anna")

			go newAnna.listenToSignal()
			go newAnna.writeStateInfo()

			newAnna.Log.WithTags(spec.Tags{L: "I", O: newAnna, T: nil, V: 10}, "booting network")
			go newAnna.Network.Boot()

			newAnna.Log.WithTags(spec.Tags{L: "I", O: newAnna, T: nil, V: 10}, "booting server")
			go newAnna.Server.Boot()

			// Block the main goroutine forever. The process is only supposed to be
			// ended by a call to Shutdown.
			select {}
		},

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var err error

			// log.
			err = newAnna.Log.SetLevels(newAnna.Flags.ControlLogLevels)
			panicOnError(err)
			err = newAnna.Log.SetObjects(newAnna.Flags.ControlLogObejcts)
			panicOnError(err)
			err = newAnna.Log.SetVerbosity(newAnna.Flags.ControlLogVerbosity)
			panicOnError(err)

			newAnna.Log.Register(newAnna.GetType())

			// text gateway.
			newTextGatewayConfig := gateway.DefaultConfig()
			newTextGatewayConfig.Log = newAnna.Log
			newTextGateway := gateway.NewGateway(newTextGatewayConfig)

			// storage.
			newStorage, err := newAnna.createStorage(newAnna.Log)
			panicOnError(err)

			// scheduler.
			newSchedulerConfig := scheduler.DefaultConfig()
			newSchedulerConfig.Log = newAnna.Log
			newSchedulerConfig.Storage = newStorage
			newScheduler, err := scheduler.NewScheduler(newSchedulerConfig)
			panicOnError(err)

			// network.
			newNetworkConfig := network.DefaultConfig()
			newNetworkConfig.Log = newAnna.Log
			newNetworkConfig.TextGateway = newTextGateway
			newNetwork, err := network.New(newNetworkConfig)
			panicOnError(err)

			// log control.
			newLogControl, err := createLogControl(newAnna.Log)
			panicOnError(err)

			// text interface.
			newTextInterface, err := createTextInterface(newAnna.Log, newScheduler, newTextGateway)
			panicOnError(err)

			// server.
			newServerConfig := server.DefaultConfig()
			newServerConfig.Addr = newAnna.Flags.Addr
			newServerConfig.Instrumentation, err = createPrometheusInstrumentation([]string{"Server"})
			panicOnError(err)
			newServerConfig.Log = newAnna.Log
			newServerConfig.LogControl = newLogControl
			newServerConfig.TextGateway = newTextGateway
			newServerConfig.TextInterface = newTextInterface
			newServer, err := server.New(newServerConfig)
			panicOnError(err)

			// Anna.
			newAnna.Network = newNetwork
			newAnna.Server = newServer
			newAnna.Storage = newStorage
		},
	}

	// Flags.
	newAnna.Cmd.PersistentFlags().StringVar(&newAnna.Flags.Addr, "addr", "127.0.0.1:9119", "host:port to bind Anna's server to")

	newAnna.Cmd.PersistentFlags().StringVar(&newAnna.Flags.ControlLogLevels, "control-log-levels", "", "set log levels for log control (e.g. E,F)")
	newAnna.Cmd.PersistentFlags().StringVar(&newAnna.Flags.ControlLogObejcts, "control-log-objects", "", "set log objects for log control (e.g. network,impulse)")
	newAnna.Cmd.PersistentFlags().IntVar(&newAnna.Flags.ControlLogVerbosity, "control-log-verbosity", 10, "set log verbosity for log control")

	newAnna.Cmd.PersistentFlags().StringVar(&newAnna.Flags.Storage, "storage", "redis", "storage type to use for persistency (e.g. memory)")
	newAnna.Cmd.PersistentFlags().StringVar(&newAnna.Flags.StorageAddr, "storage-addr", "127.0.0.1:6379", "host:port to connect to storage")

	return newAnna, nil
}

type anna struct {
	Config

	Cmd          *cobra.Command
	BootOnce     sync.Once
	ID           spec.ObjectID
	ShutdownOnce sync.Once
	Type         spec.ObjectType
}

func (a *anna) Boot() {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call Boot")

	a.BootOnce.Do(func() {
		// init
		a.Cmd.AddCommand(a.InitVersionCmd())

		// execute
		a.Cmd.Execute()
	})
}

func (a *anna) Shutdown() {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call Shutdown")

	a.ShutdownOnce.Do(func() {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "shutting down network")
			a.Network.Shutdown()
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "shutting down server")
			a.Server.Shutdown()
			wg.Done()
		}()

		wg.Wait()

		a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "shutting down Anna")
		os.Exit(0)
	})
}

func main() {
	newAnna, err := New(DefaultConfig())
	panicOnError(err)

	newAnna.Boot()
}
