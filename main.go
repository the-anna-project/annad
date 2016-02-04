package main

import (
	"os"

	"github.com/spf13/pflag"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/factory/server"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/server"
	"github.com/xh3b4sd/anna/spec"
)

var (
	stateReader string
	stateWriter string

	verbosity int
)

func init() {
	pflag.StringVar(&stateReader, "state-reader", string(common.StateType.FSReader), "where to read state from")
	pflag.StringVar(&stateWriter, "state-writer", string(common.StateType.FSWriter), "where to write state to")

	pflag.IntVarP(&verbosity, "verbosity", "v", 12, "verbosity of the logger: 0 - 12")

	pflag.Parse()
}

func main() {
	//
	// create dependencies
	//
	newFactoryGateway := gateway.NewGateway()
	newLogConfig := log.DefaultConfig()
	newLogConfig.Verbosity = verbosity
	newLog := log.NewLog(newLogConfig)
	newTextGateway := gateway.NewGateway()

	newLog.V(9).Infof("hello, I am Anna")

	//
	// create factory
	//
	newLog.V(9).Infof("creating factory")
	newFactoryClientConfig := factoryclient.DefaultConfig()
	newFactoryClientConfig.FactoryGateway = newFactoryGateway
	newFactoryClientConfig.Log = newLog
	newFactoryGatewayClient := factoryclient.NewClient(newFactoryClientConfig)
	newFactoryServerConfig := factoryserver.DefaultConfig()
	newFactoryServerConfig.FactoryClient = newFactoryGatewayClient
	newFactoryServerConfig.FactoryGateway = newFactoryGateway
	newFactoryServerConfig.Log = newLog
	newFactoryServerConfig.StateReader = spec.StateType(stateReader)
	newFactoryServerConfig.StateWriter = spec.StateType(stateWriter)
	newFactoryServerConfig.TextGateway = newTextGateway
	newFactoryServer := factoryserver.NewServer(newFactoryServerConfig)

	//
	// create core
	//
	newLog.V(9).Infof("creating core")
	newCore, err := newFactoryServer.NewCore()
	if err != nil {
		newLog.V(3).Errorf("%#v", maskAny(err))
		os.Exit(0)
	}
	err = newCore.GetState().Read()
	if err != nil {
		newLog.V(3).Errorf("%#v", maskAny(err))
		os.Exit(0)
	}
	go newCore.Boot()

	//
	// create server
	//
	newLog.V(9).Infof("creating server")
	newServerConfig := server.DefaultConfig()
	newServerConfig.Log = newLog
	newServerConfig.TextGateway = newTextGateway
	newServer := server.NewServer(newServerConfig)
	newServer.Listen()
}
