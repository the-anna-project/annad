package main

import (
	"github.com/juju/errgo"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/server"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

func main() {
	//
	// configure dependencies
	//
	newLog := log.NewLog(log.DefaultConfig())
	newLog.V(9).Infof("hello, I am Anna")
	newTextGateway := gateway.NewGateway()

	newInitPairConfig := initPairConfig{
		TextGateway: newTextGateway,
		Log:         newLog,
	}

	//
	// initialize core
	//
	impulseInitPair := mustGetImpulseInitPair(newInitPairConfig)
	firstNeuronInitPair := mustGetFirstNeuronInitPair(newInitPairConfig)
	jobNeuronInitPair := mustGetJobNeuronInitPair(newInitPairConfig)
	characterNeuronInitPair := mustGetCharacterNeuronInitPair(newInitPairConfig)
	networkInitPair := mustGetNetworkInitPair(newInitPairConfig)
	coreInitPair := mustGetCoreInitPair(newInitPairConfig)

	initPairs := []initPair{
		impulseInitPair,
		firstNeuronInitPair,
		jobNeuronInitPair,
		characterNeuronInitPair,
		networkInitPair,
		coreInitPair,
	}

	for _, ip := range initPairs {
		ip.State.SetImpulse(common.MustObjectToImpulse(impulseInitPair.Object))
		ip.State.SetBytes(common.ImpulseIDKey, []byte(impulseInitPair.Object.GetObjectID()))

		ip.State.SetNeuron(common.MustObjectToNeuron(firstNeuronInitPair.Object))
		ip.State.SetBytes(common.FirstNeuronIDKey, []byte(firstNeuronInitPair.Object.GetObjectID()))

		ip.State.SetNeuron(common.MustObjectToNeuron(jobNeuronInitPair.Object))
		ip.State.SetBytes(common.JobNeuronIDKey, []byte(jobNeuronInitPair.Object.GetObjectID()))

		ip.State.SetNeuron(common.MustObjectToNeuron(characterNeuronInitPair.Object))
		ip.State.SetBytes(common.CharacterNeuronIDKey, []byte(characterNeuronInitPair.Object.GetObjectID()))

		ip.State.SetNetwork(common.MustObjectToNetwork(networkInitPair.Object))
		ip.State.SetBytes(common.NetworkIDKey, []byte(networkInitPair.Object.GetObjectID()))

		ip.State.SetCore(common.MustObjectToCore(coreInitPair.Object))
		ip.State.SetBytes(common.CoreIDKey, []byte(coreInitPair.Object.GetObjectID()))

		ip.Object.SetState(common.InitStateKey, ip.State)
	}

	newLog.V(9).Infof("booting core")
	go common.MustObjectToCore(coreInitPair.Object).Boot()

	//
	// initialize server
	//
	newServerConfig := server.DefaultConfig()
	newServerConfig.Log = newLog
	newServerConfig.TextGateway = newTextGateway
	newServer := server.NewServer(newServerConfig)

	newLog.V(9).Infof("starting server")
	go newServer.Listen()

	//
	// initialization finished
	//
	for {
	}
}
