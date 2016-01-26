package main

import (
	"github.com/juju/errgo"

	"github.com/xh3b4sd/anna/core"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/network"
	"github.com/xh3b4sd/anna/neuron"
	"github.com/xh3b4sd/anna/server"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

func main() {
	//
	// deps
	//
	newLog := log.NewLog(log.DefaultConfig())
	newLog.V(7).Infof("%s", "hello, I am Anna")

	newTextGateway := gateway.NewGateway()

	//
	// core
	//
	newImpulseConfig := impulse.DefaultConfig()
	newImpulseConfig.Log = newLog
	newImpulse := impulse.NewImpulse(newImpulseConfig)

	newFirstConfig := neuron.DefaultFirstNeuronConfig()
	newFirstConfig.Log = newLog
	newFirst := neuron.NewFirstNeuron(newFirstConfig)

	newJobConfig := neuron.DefaultJobNeuronConfig()
	newJobConfig.Log = newLog
	newJob := neuron.NewJobNeuron(newJobConfig)

	newCharacterConfig := neuron.DefaultCharacterNeuronConfig()
	newCharacterConfig.Log = newLog
	newCharacter := neuron.NewCharacterNeuron(newCharacterConfig)

	newNetworkConfig := network.DefaultConfig()
	newNetworkConfig.Log = newLog
	newNetwork := network.NewNetwork(newNetworkConfig)

	newCoreConfig := core.DefaultConfig()
	newCoreConfig.Log = newLog
	newCoreConfig.TextGateway = newTextGateway
	newCore := core.NewCore(newCoreConfig)

	impulseState, err := newImpulse.GetState("default")
	if err != nil {
		newLog.V(1).Errorf("%#v", maskAny(err))
	}
	impulseState = impulseState.Copy()
	impulseState.SetImpulse(newImpulse)
	impulseState.SetBytes("impulse-id", []byte(newImpulse.GetObjectID()))
	impulseState.SetNeuron(newFirst)
	impulseState.SetBytes("first-id", []byte(newFirst.GetObjectID()))
	impulseState.SetNeuron(newJob)
	impulseState.SetBytes("job-id", []byte(newJob.GetObjectID()))
	impulseState.SetNeuron(newCharacter)
	impulseState.SetBytes("character-id", []byte(newCharacter.GetObjectID()))
	impulseState.SetNetwork(newNetwork)
	impulseState.SetBytes("network-id", []byte(newNetwork.GetObjectID()))
	impulseState.SetCore(newCore)
	impulseState.SetBytes("core-id", []byte(newCore.GetObjectID()))
	newImpulse.SetState("init", impulseState)

	firstState, err := newFirst.GetState("default")
	if err != nil {
		newLog.V(1).Errorf("%#v", maskAny(err))
	}
	firstState = firstState.Copy()
	firstState.SetImpulse(newImpulse)
	firstState.SetBytes("impulse-id", []byte(newImpulse.GetObjectID()))
	firstState.SetNeuron(newFirst)
	firstState.SetBytes("first-id", []byte(newFirst.GetObjectID()))
	firstState.SetNeuron(newJob)
	firstState.SetBytes("job-id", []byte(newJob.GetObjectID()))
	firstState.SetNeuron(newCharacter)
	firstState.SetBytes("character-id", []byte(newCharacter.GetObjectID()))
	firstState.SetNetwork(newNetwork)
	firstState.SetBytes("network-id", []byte(newNetwork.GetObjectID()))
	firstState.SetCore(newCore)
	firstState.SetBytes("core-id", []byte(newCore.GetObjectID()))
	newFirst.SetState("init", firstState)

	jobState, err := newJob.GetState("default")
	if err != nil {
		newLog.V(1).Errorf("%#v", maskAny(err))
	}
	jobState = jobState.Copy()
	jobState.SetImpulse(newImpulse)
	jobState.SetBytes("impulse-id", []byte(newImpulse.GetObjectID()))
	jobState.SetNeuron(newFirst)
	jobState.SetBytes("first-id", []byte(newFirst.GetObjectID()))
	jobState.SetNeuron(newJob)
	jobState.SetBytes("job-id", []byte(newJob.GetObjectID()))
	jobState.SetNeuron(newCharacter)
	jobState.SetBytes("character-id", []byte(newCharacter.GetObjectID()))
	jobState.SetNetwork(newNetwork)
	jobState.SetBytes("network-id", []byte(newNetwork.GetObjectID()))
	jobState.SetCore(newCore)
	jobState.SetBytes("core-id", []byte(newCore.GetObjectID()))
	newJob.SetState("init", jobState)

	characterState, err := newCharacter.GetState("default")
	if err != nil {
		newLog.V(1).Errorf("%#v", maskAny(err))
	}
	characterState = characterState.Copy()
	characterState.SetImpulse(newImpulse)
	characterState.SetBytes("impulse-id", []byte(newImpulse.GetObjectID()))
	characterState.SetNeuron(newFirst)
	characterState.SetBytes("first-id", []byte(newFirst.GetObjectID()))
	characterState.SetNeuron(newJob)
	characterState.SetBytes("job-id", []byte(newJob.GetObjectID()))
	characterState.SetNeuron(newCharacter)
	characterState.SetBytes("character-id", []byte(newCharacter.GetObjectID()))
	characterState.SetNetwork(newNetwork)
	characterState.SetBytes("network-id", []byte(newNetwork.GetObjectID()))
	characterState.SetCore(newCore)
	characterState.SetBytes("core-id", []byte(newCore.GetObjectID()))
	newCharacter.SetState("init", characterState)

	networkState, err := newNetwork.GetState("default")
	if err != nil {
		newLog.V(1).Errorf("%#v", maskAny(err))
	}
	networkState = networkState.Copy()
	networkState.SetImpulse(newImpulse)
	networkState.SetBytes("impulse-id", []byte(newImpulse.GetObjectID()))
	networkState.SetNeuron(newFirst)
	networkState.SetBytes("first-id", []byte(newFirst.GetObjectID()))
	networkState.SetNeuron(newJob)
	networkState.SetBytes("job-id", []byte(newJob.GetObjectID()))
	networkState.SetNeuron(newCharacter)
	networkState.SetBytes("character-id", []byte(newCharacter.GetObjectID()))
	networkState.SetNetwork(newNetwork)
	networkState.SetBytes("network-id", []byte(newNetwork.GetObjectID()))
	networkState.SetCore(newCore)
	networkState.SetBytes("core-id", []byte(newCore.GetObjectID()))
	newNetwork.SetState("init", networkState)

	coreState, err := newCore.GetState("default")
	if err != nil {
		newLog.V(1).Errorf("%#v", maskAny(err))
	}
	coreState = coreState.Copy()
	coreState.SetImpulse(newImpulse)
	coreState.SetBytes("impulse-id", []byte(newImpulse.GetObjectID()))
	coreState.SetNeuron(newFirst)
	coreState.SetBytes("first-id", []byte(newFirst.GetObjectID()))
	coreState.SetNeuron(newJob)
	coreState.SetBytes("job-id", []byte(newJob.GetObjectID()))
	coreState.SetNeuron(newCharacter)
	coreState.SetBytes("character-id", []byte(newCharacter.GetObjectID()))
	coreState.SetNetwork(newNetwork)
	coreState.SetBytes("network-id", []byte(newNetwork.GetObjectID()))
	coreState.SetCore(newCore)
	coreState.SetBytes("core-id", []byte(newCore.GetObjectID()))
	newCore.SetState("init", coreState)

	newLog.V(7).Infof("%s", "booting core")
	go newCore.Boot()

	//
	// server
	//
	newServerConfig := server.DefaultConfig()
	newServerConfig.Log = newLog
	newServerConfig.TextGateway = newTextGateway
	newServer := server.NewServer(newServerConfig)

	newLog.V(7).Infof("%s", "starting server")
	go newServer.Listen()

	for {
	}
}
