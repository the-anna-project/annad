package main

import (
	"github.com/juju/errgo"

	"github.com/xh3b4sd/anna/common"
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

	impulseState, err := newImpulse.GetState(common.DefaultStateKey)
	if err != nil {
		newLog.V(1).Errorf("%#v", maskAny(err))
	}
	impulseState = impulseState.Copy()
	impulseState.SetImpulse(newImpulse)
	impulseState.SetBytes(common.ImpulseIDKey, []byte(newImpulse.GetObjectID()))
	impulseState.SetNeuron(newFirst)
	impulseState.SetBytes(common.FirstIDKey, []byte(newFirst.GetObjectID()))
	impulseState.SetNeuron(newJob)
	impulseState.SetBytes(common.JobIDKey, []byte(newJob.GetObjectID()))
	impulseState.SetNeuron(newCharacter)
	impulseState.SetBytes(common.CharacterIDKey, []byte(newCharacter.GetObjectID()))
	impulseState.SetNetwork(newNetwork)
	impulseState.SetBytes(common.NetworkIDKey, []byte(newNetwork.GetObjectID()))
	impulseState.SetCore(newCore)
	impulseState.SetBytes(common.CoreIDKey, []byte(newCore.GetObjectID()))
	newImpulse.SetState(common.InitStateKey, impulseState)

	firstState, err := newFirst.GetState(common.DefaultStateKey)
	if err != nil {
		newLog.V(1).Errorf("%#v", maskAny(err))
	}
	firstState = firstState.Copy()
	firstState.SetImpulse(newImpulse)
	firstState.SetBytes(common.ImpulseIDKey, []byte(newImpulse.GetObjectID()))
	firstState.SetNeuron(newFirst)
	firstState.SetBytes(common.FirstIDKey, []byte(newFirst.GetObjectID()))
	firstState.SetNeuron(newJob)
	firstState.SetBytes(common.JobIDKey, []byte(newJob.GetObjectID()))
	firstState.SetNeuron(newCharacter)
	firstState.SetBytes(common.CharacterIDKey, []byte(newCharacter.GetObjectID()))
	firstState.SetNetwork(newNetwork)
	firstState.SetBytes(common.NetworkIDKey, []byte(newNetwork.GetObjectID()))
	firstState.SetCore(newCore)
	firstState.SetBytes(common.CoreIDKey, []byte(newCore.GetObjectID()))
	newFirst.SetState(common.InitStateKey, firstState)

	jobState, err := newJob.GetState(common.DefaultStateKey)
	if err != nil {
		newLog.V(1).Errorf("%#v", maskAny(err))
	}
	jobState = jobState.Copy()
	jobState.SetImpulse(newImpulse)
	jobState.SetBytes(common.ImpulseIDKey, []byte(newImpulse.GetObjectID()))
	jobState.SetNeuron(newFirst)
	jobState.SetBytes(common.FirstIDKey, []byte(newFirst.GetObjectID()))
	jobState.SetNeuron(newJob)
	jobState.SetBytes(common.JobIDKey, []byte(newJob.GetObjectID()))
	jobState.SetNeuron(newCharacter)
	jobState.SetBytes(common.CharacterIDKey, []byte(newCharacter.GetObjectID()))
	jobState.SetNetwork(newNetwork)
	jobState.SetBytes(common.NetworkIDKey, []byte(newNetwork.GetObjectID()))
	jobState.SetCore(newCore)
	jobState.SetBytes(common.CoreIDKey, []byte(newCore.GetObjectID()))
	newJob.SetState(common.InitStateKey, jobState)

	characterState, err := newCharacter.GetState(common.DefaultStateKey)
	if err != nil {
		newLog.V(1).Errorf("%#v", maskAny(err))
	}
	characterState = characterState.Copy()
	characterState.SetImpulse(newImpulse)
	characterState.SetBytes(common.ImpulseIDKey, []byte(newImpulse.GetObjectID()))
	characterState.SetNeuron(newFirst)
	characterState.SetBytes(common.FirstIDKey, []byte(newFirst.GetObjectID()))
	characterState.SetNeuron(newJob)
	characterState.SetBytes(common.JobIDKey, []byte(newJob.GetObjectID()))
	characterState.SetNeuron(newCharacter)
	characterState.SetBytes(common.CharacterIDKey, []byte(newCharacter.GetObjectID()))
	characterState.SetNetwork(newNetwork)
	characterState.SetBytes(common.NetworkIDKey, []byte(newNetwork.GetObjectID()))
	characterState.SetCore(newCore)
	characterState.SetBytes(common.CoreIDKey, []byte(newCore.GetObjectID()))
	newCharacter.SetState(common.InitStateKey, characterState)

	networkState, err := newNetwork.GetState(common.DefaultStateKey)
	if err != nil {
		newLog.V(1).Errorf("%#v", maskAny(err))
	}
	networkState = networkState.Copy()
	networkState.SetImpulse(newImpulse)
	networkState.SetBytes(common.ImpulseIDKey, []byte(newImpulse.GetObjectID()))
	networkState.SetNeuron(newFirst)
	networkState.SetBytes(common.FirstIDKey, []byte(newFirst.GetObjectID()))
	networkState.SetNeuron(newJob)
	networkState.SetBytes(common.JobIDKey, []byte(newJob.GetObjectID()))
	networkState.SetNeuron(newCharacter)
	networkState.SetBytes(common.CharacterIDKey, []byte(newCharacter.GetObjectID()))
	networkState.SetNetwork(newNetwork)
	networkState.SetBytes(common.NetworkIDKey, []byte(newNetwork.GetObjectID()))
	networkState.SetCore(newCore)
	networkState.SetBytes(common.CoreIDKey, []byte(newCore.GetObjectID()))
	newNetwork.SetState(common.InitStateKey, networkState)

	coreState, err := newCore.GetState(common.DefaultStateKey)
	if err != nil {
		newLog.V(1).Errorf("%#v", maskAny(err))
	}
	coreState = coreState.Copy()
	coreState.SetImpulse(newImpulse)
	coreState.SetBytes(common.ImpulseIDKey, []byte(newImpulse.GetObjectID()))
	coreState.SetNeuron(newFirst)
	coreState.SetBytes(common.FirstIDKey, []byte(newFirst.GetObjectID()))
	coreState.SetNeuron(newJob)
	coreState.SetBytes(common.JobIDKey, []byte(newJob.GetObjectID()))
	coreState.SetNeuron(newCharacter)
	coreState.SetBytes(common.CharacterIDKey, []byte(newCharacter.GetObjectID()))
	coreState.SetNetwork(newNetwork)
	coreState.SetBytes(common.NetworkIDKey, []byte(newNetwork.GetObjectID()))
	coreState.SetCore(newCore)
	coreState.SetBytes(common.CoreIDKey, []byte(newCore.GetObjectID()))
	newCore.SetState(common.InitStateKey, coreState)

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
