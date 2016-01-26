package main

import (
	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/core"
	gatewayspec "github.com/xh3b4sd/anna/gateway/spec"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/network"
	"github.com/xh3b4sd/anna/neuron"
	"github.com/xh3b4sd/anna/spec"
)

type initPair struct {
	Object spec.Object
	State  spec.State
}

type initPairConfig struct {
	TextGateway gatewayspec.Gateway
	Log         spec.Log
}

func mustGetImpulseInitPair(config initPairConfig) initPair {
	newImpulseConfig := impulse.DefaultConfig()
	newImpulseConfig.Log = config.Log
	newImpulse := impulse.NewImpulse(newImpulseConfig)
	newImpulseState, err := newImpulse.GetState(common.DefaultStateKey)
	if err != nil {
		panic(err)
	}
	newImpulseState = newImpulseState.Copy()

	return initPair{Object: newImpulse, State: newImpulseState}
}

func mustGetFirstNeuronInitPair(config initPairConfig) initPair {
	newFirstNeuronConfig := neuron.DefaultFirstNeuronConfig()
	newFirstNeuronConfig.Log = config.Log
	newFirstNeuron := neuron.NewFirstNeuron(newFirstNeuronConfig)
	newFirstNeuronState, err := newFirstNeuron.GetState(common.DefaultStateKey)
	if err != nil {
		panic(err)
	}
	newFirstNeuronState = newFirstNeuronState.Copy()

	return initPair{Object: newFirstNeuron, State: newFirstNeuronState}
}

func mustGetJobNeuronInitPair(config initPairConfig) initPair {
	newJobNeuronConfig := neuron.DefaultJobNeuronConfig()
	newJobNeuronConfig.Log = config.Log
	newJobNeuron := neuron.NewJobNeuron(newJobNeuronConfig)
	newJobNeuronState, err := newJobNeuron.GetState(common.DefaultStateKey)
	if err != nil {
		panic(err)
	}
	newJobNeuronState = newJobNeuronState.Copy()

	return initPair{Object: newJobNeuron, State: newJobNeuronState}
}

func mustGetCharacterNeuronInitPair(config initPairConfig) initPair {
	newCharacterNeuronConfig := neuron.DefaultCharacterNeuronConfig()
	newCharacterNeuronConfig.Log = config.Log
	newCharacterNeuron := neuron.NewCharacterNeuron(newCharacterNeuronConfig)
	newCharacterNeuronState, err := newCharacterNeuron.GetState(common.DefaultStateKey)
	if err != nil {
		panic(err)
	}
	newCharacterNeuronState = newCharacterNeuronState.Copy()

	return initPair{Object: newCharacterNeuron, State: newCharacterNeuronState}
}

func mustGetNetworkInitPair(config initPairConfig) initPair {
	newNetworkConfig := network.DefaultConfig()
	newNetworkConfig.Log = config.Log
	newNetwork := network.NewNetwork(newNetworkConfig)
	newNetworkState, err := newNetwork.GetState(common.DefaultStateKey)
	if err != nil {
		panic(err)
	}
	newNetworkState = newNetworkState.Copy()

	return initPair{Object: newNetwork, State: newNetworkState}
}

func mustGetCoreInitPair(config initPairConfig) initPair {
	newCoreConfig := core.DefaultConfig()
	newCoreConfig.Log = config.Log
	newCoreConfig.TextGateway = config.TextGateway
	newCore := core.NewCore(newCoreConfig)
	newCoreState, err := newCore.GetState(common.DefaultStateKey)
	if err != nil {
		panic(err)
	}
	newCoreState = newCoreState.Copy()

	return initPair{Object: newCore, State: newCoreState}
}
