package main

import (
	"github.com/xh3b4sd/anna/core"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/server"
)

func main() {
	newLog := log.NewLog(log.DefaultConfig())
	newLog.V(6).Infof("%s", "hello, I am Anna")

	textGateway := gateway.NewGateway()

	newLog.V(6).Infof("%s", "booting core")
	newCoreConfig := core.DefaultConfig()
	newCoreConfig.TextGateway = textGateway
	newCore := core.NewCore(newCoreConfig)
	go newCore.Boot()

	newLog.V(6).Infof("%s", "starting server")
	newServerConfig := server.DefaultConfig()
	newServerConfig.TextGateway = textGateway
	newServer := server.NewServer(newServerConfig)
	go newServer.Listen()

	for {
	}
}
