package main

import (
	"os"
	"sync"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/core"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/factory/server"
	"github.com/xh3b4sd/anna/file-system/os"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/server"
	"github.com/xh3b4sd/anna/spec"
)

const (
	objectTypeAnna spec.ObjectType = "anna"
)

var (
	globalFlags struct {
		ControlLogLevels    string
		ControlLogObejcts   string
		ControlLogVerbosity int

		Addr string
	}

	annaCmd = &cobra.Command{
		Use:   "anna",
		Short: "Anna, Artificial Neural Network Aspiration, aims to be self-learning and self-improving software. For more information see https://github.com/xh3b4sd/anna.",
		Long:  "Anna, Artificial Neural Network Aspiration, aims to be self-learning and self-improving software. For more information see https://github.com/xh3b4sd/anna.",
		Run:   mainRun,
	}

	// Version is the project version. It is given via buildflags that inject the
	// commit hash.
	version string
)

func init() {
	annaCmd.PersistentFlags().StringVar(&globalFlags.ControlLogLevels, "control-log-levels", "", "set log levels for log control (e.g. E,F)")
	annaCmd.PersistentFlags().StringVar(&globalFlags.ControlLogObejcts, "control-log-objects", "", "set log objects for log control (e.g. core,network)")
	annaCmd.PersistentFlags().IntVar(&globalFlags.ControlLogVerbosity, "control-log-verbosity", 10, "set log verbosity for log control")

	annaCmd.PersistentFlags().StringVar(&globalFlags.Addr, "addr", "127.0.0.1:9119", "host:port to bind Anna's server to")
}

type annaConfig struct {
	Core spec.Core

	FactoryServer spec.Factory

	Log spec.Log

	Server spec.Server
}

func defaultAnnaConfig() annaConfig {
	newConfig := annaConfig{
		Core:          core.NewCore(core.DefaultConfig()),
		FactoryServer: factoryserver.NewFactory(factoryserver.DefaultConfig()),
		Log:           log.NewLog(log.DefaultConfig()),
		Server:        server.NewServer(server.DefaultConfig()),
	}

	return newConfig
}

func newAnna(config annaConfig) spec.Anna {
	newAnna := &anna{
		annaConfig: config,
		ID:         id.NewObjectID(id.Hex128),
		Mutex:      sync.Mutex{},
		Type:       spec.ObjectType(objectTypeAnna),
	}

	newAnna.Log.Register(newAnna.GetType())

	return newAnna
}

// mainObject is basically only to have an object that provides proper
// identifyable logging.
type anna struct {
	annaConfig

	ID spec.ObjectID

	Mutex sync.Mutex

	Type spec.ObjectType
}

func (a *anna) Boot() {
	a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "hello, I am Anna")

	go a.listenToSignal()

	a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "booting factory")
	go a.FactoryServer.Boot()

	a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "booting core")
	go a.Core.Boot()

	a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "booting server")
	a.Server.Boot()
}

func (a *anna) Shutdown() {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 13}, "call Shutdown")

	go a.Core.Shutdown()
	go a.FactoryServer.Shutdown()

	a.Log.WithTags(spec.Tags{L: "I", O: a, T: nil, V: 10}, "shutting down")
	os.Exit(0)
}

func mainRun(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Help()
		os.Exit(1)
	}

	//
	// create dependencies
	//
	newLog := log.NewLog(log.DefaultConfig())
	newLog.SetLevels(globalFlags.ControlLogLevels)
	newLog.SetObjects(globalFlags.ControlLogObejcts)
	newLog.SetVerbosity(globalFlags.ControlLogVerbosity)

	newFactoryGatewayConfig := gateway.DefaultConfig()
	newFactoryGatewayConfig.Log = newLog
	newFactoryGateway := gateway.NewGateway(newFactoryGatewayConfig)

	newTextGatewayConfig := gateway.DefaultConfig()
	newTextGatewayConfig.Log = newLog
	newTextGateway := gateway.NewGateway(newTextGatewayConfig)

	newOSFileSystem := osfilesystem.NewFileSystem()

	//
	// create factory
	//
	newFactoryClientConfig := factoryclient.DefaultConfig()
	newFactoryClientConfig.FactoryGateway = newFactoryGateway
	newFactoryClientConfig.Log = newLog
	newFactoryGatewayClient := factoryclient.NewFactory(newFactoryClientConfig)
	newFactoryServerConfig := factoryserver.DefaultConfig()
	newFactoryServerConfig.FactoryClient = newFactoryGatewayClient
	newFactoryServerConfig.FactoryGateway = newFactoryGateway
	newFactoryServerConfig.FileSystem = newOSFileSystem
	newFactoryServerConfig.Log = newLog
	newFactoryServerConfig.TextGateway = newTextGateway
	newFactoryServer := factoryserver.NewFactory(newFactoryServerConfig)

	//
	// create core
	//
	newCore, err := newFactoryServer.NewCore()
	if err != nil {
		panic(err)
	}

	//
	// create server
	//
	newServerConfig := server.DefaultConfig()
	newServerConfig.Addr = globalFlags.Addr
	newServerConfig.Log = newLog
	newServerConfig.TextGateway = newTextGateway
	newServer := server.NewServer(newServerConfig)

	//
	// create anna
	//
	newAnnaConfig := defaultAnnaConfig()
	newAnnaConfig.Core = newCore
	newAnnaConfig.FactoryServer = newFactoryServer
	newAnnaConfig.Log = newLog
	newAnnaConfig.Server = newServer
	a := newAnna(newAnnaConfig)

	a.Boot()
}

func main() {
	annaCmd.AddCommand(versionCmd)

	annaCmd.Execute()
}
