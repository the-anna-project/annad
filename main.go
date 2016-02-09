package main

import (
	"github.com/spf13/pflag"

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
	stateReader string
	stateWriter string

	logTags struct {
		L string
		O string
		V int
	}
)

func init() {
	pflag.StringVar(&stateReader, "state-reader", string(common.StateType.FSReader), "where to read state from")
	pflag.StringVar(&stateWriter, "state-writer", string(common.StateType.FSWriter), "where to write state to")

	pflag.StringVar(&logTags.L, "log-tag-l", "", "levels tags of the logger: comma separated")
	pflag.StringVar(&logTags.O, "log-tag-o", "", "objects tag of the logger: comma separated")
	pflag.IntVar(&logTags.V, "log-tag-v", 10, "verbosity tag of the logger: 0 - 15")

	pflag.Parse()
}

func main() {
	//
	// create main object
	//
	m := mainO{}

	//
	// create dependencies
	//
	newFactoryGateway := gateway.NewGateway()
	newLog := log.NewLog(log.DefaultConfig())
	newLog.SetLevels(logTags.L)
	newLog.SetObjects(logTags.O)
	newLog.SetVerbosity(logTags.V)
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
	newFactoryServerConfig.StateReader = spec.StateType(stateReader)
	newFactoryServerConfig.StateWriter = spec.StateType(stateWriter)
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
	newServerConfig.Log = newLog
	newServerConfig.TextGateway = newTextGateway
	newServer := server.NewServer(newServerConfig)
	newServer.Listen()
}
