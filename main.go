package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/factory/server"
	"github.com/xh3b4sd/anna/file-system/real"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/server"
	"github.com/xh3b4sd/anna/spec"
)

var (
	globalFlags struct {
		ControlLogLevels    string
		ControlLogObejcts   string
		ControlLogVerbosity int

		Host string

		StateReader string
		StateWriter string
	}

	mainCmd = &cobra.Command{
		Use:   "anna",
		Short: "artificial neural network aspiration",
		Long:  "artificial neural network aspiration",
		Run:   mainRun,
	}
)

func init() {
	mainCmd.PersistentFlags().StringVar(&globalFlags.ControlLogLevels, "control-log-levels", "", "set log levels for log control (e.g. E,F)")
	mainCmd.PersistentFlags().StringVar(&globalFlags.ControlLogObejcts, "control-log-objects", "", "set log objects for log control (e.g. core,network)")
	mainCmd.PersistentFlags().IntVar(&globalFlags.ControlLogVerbosity, "control-log-verbosity", 10, "set log verbosity for log control")

	mainCmd.PersistentFlags().StringVar(&globalFlags.Host, "host", "127.0.0.1:9119", "host:port to bind Anna's server to")

	mainCmd.PersistentFlags().StringVar(&globalFlags.StateReader, "state-reader", string(common.StateType.FSReader), "where to read state from")
	mainCmd.PersistentFlags().StringVar(&globalFlags.StateWriter, "state-writer", string(common.StateType.FSWriter), "where to write state to")
}

func mainRun(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		cmd.Help()
		os.Exit(1)
	}

	//
	// create main object
	//
	m := mainO{}

	//
	// create dependencies
	//
	newFactoryGateway := gateway.NewGateway()
	newLog := log.NewLog(log.DefaultConfig())
	newLog.SetLevels(globalFlags.ControlLogLevels)
	newLog.SetObjects(globalFlags.ControlLogObejcts)
	newLog.SetVerbosity(globalFlags.ControlLogVerbosity)
	newTextGateway := gateway.NewGateway()
	newFileSystemReal := filesystemreal.NewFileSystem()

	newLog.WithTags(spec.Tags{L: "I", O: m, T: nil, V: 10}, "hello, I am Anna")

	//
	// create factory
	//
	newLog.WithTags(spec.Tags{L: "I", O: m, T: nil, V: 10}, "creating factory")

	newFactoryClientConfig := factoryclient.DefaultConfig()
	newFactoryClientConfig.FactoryGateway = newFactoryGateway
	newFactoryClientConfig.Log = newLog
	newFactoryGatewayClient := factoryclient.NewFactory(newFactoryClientConfig)
	newFactoryServerConfig := factoryserver.DefaultConfig()
	newFactoryServerConfig.FactoryClient = newFactoryGatewayClient
	newFactoryServerConfig.FactoryGateway = newFactoryGateway
	newFactoryServerConfig.FileSystem = newFileSystemReal
	newFactoryServerConfig.Log = newLog
	newFactoryServerConfig.StateReader = spec.StateType(globalFlags.StateReader)
	newFactoryServerConfig.StateWriter = spec.StateType(globalFlags.StateWriter)
	newFactoryServerConfig.TextGateway = newTextGateway
	newFactoryServer := factoryserver.NewFactory(newFactoryServerConfig)

	//
	// create core
	//
	newLog.WithTags(spec.Tags{L: "I", O: m, T: nil, V: 10}, "creating core")

	newCore, err := newFactoryServer.NewCore()
	if err != nil {
		newLog.WithTags(spec.Tags{L: "F", O: m, T: nil, V: 1}, "%#v", maskAny(err))
	}
	go newCore.Boot()

	//
	// create server
	//
	newLog.WithTags(spec.Tags{L: "I", O: m, T: nil, V: 10}, "creating server")

	newServerConfig := server.DefaultConfig()
	newServerConfig.Host = globalFlags.Host
	newServerConfig.Log = newLog
	newServerConfig.TextGateway = newTextGateway
	newServer := server.NewServer(newServerConfig)
	newServer.Listen()
}

func main() {
	mainCmd.Execute()
}
