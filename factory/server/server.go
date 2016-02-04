package factoryserver

import (
	"time"

	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/core"
	"github.com/xh3b4sd/anna/factory/client"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/gateway/spec"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/network"
	"github.com/xh3b4sd/anna/neuron"
	"github.com/xh3b4sd/anna/spec"
	"github.com/xh3b4sd/anna/state"
)

type Config struct {
	FactoryClient  spec.Factory
	FactoryGateway gatewayspec.Gateway
	Log            spec.Log
	StateReader    spec.StateType
	StateWriter    spec.StateType
	TextGateway    gatewayspec.Gateway
}

func DefaultConfig() Config {
	newConfig := Config{
		FactoryClient:  factoryclient.NewClient(factoryclient.DefaultConfig()),
		FactoryGateway: gateway.NewGateway(),
		Log:            log.NewLog(log.DefaultConfig()),
		StateReader:    common.StateType.FSReader,
		StateWriter:    common.StateType.FSWriter,
		TextGateway:    gateway.NewGateway(),
	}

	return newConfig
}

func NewServer(config Config) spec.Factory {
	newFactory := &server{
		Config: config,
	}

	go newFactory.listenToGateway()

	return newFactory
}

type server struct {
	Config
}

func (s *server) listenToGateway() {
	s.Log.V(11).Debugf("call Factory.listenToGateway")

	for {
		newSignal, err := s.FactoryGateway.ReceiveSignal()
		if gateway.IsGatewayClosed(err) {
			s.Log.V(6).Warnf("gateway is closed")
			time.Sleep(1 * time.Second)
			continue
		} else if err != nil {
			s.Log.V(3).Errorf("%#v", maskAny(err))
			continue
		}
		s.Log.V(12).Debugf("factory received new signal '%s'", newSignal.GetID())

		responder, err := newSignal.GetResponder()
		if gateway.IsSignalCanceled(err) {
			s.Log.V(6).Warnf("signal is canceled")
			continue
		} else if err != nil {
			s.Log.V(3).Errorf("%#v", maskAny(err))
			continue
		}

		request, err := newSignal.GetBytes("request")
		if err != nil {
			s.Log.V(3).Errorf("%#v", maskAny(err))
			newSignal.SetError(maskAny(err))
			responder <- newSignal
			continue
		}

		var response interface{}
		switch spec.ObjectType(request) {
		case common.ObjectType.Core:
			response, err = s.NewCore()
		case common.ObjectType.Impulse:
			response, err = s.NewImpulse()
		case common.ObjectType.CharacterNeuron:
			response, err = s.NewCharacterNeuron()
		case common.ObjectType.FirstNeuron:
			response, err = s.NewFirstNeuron()
		case common.ObjectType.JobNeuron:
			response, err = s.NewJobNeuron()
		case common.ObjectType.Network:
			response, err = s.NewNetwork()
		case common.ObjectType.State:
			stateObjectType, err := newSignal.GetBytes("state-object-type")
			if err != nil {
				s.Log.V(3).Errorf("%#v", maskAny(err))
				newSignal.SetError(maskAny(err))
				responder <- newSignal
				continue
			}
			response, err = s.NewState(spec.ObjectType(stateObjectType))
		default:
			s.Log.V(3).Errorf("%#v", maskAnyf(invalidFactoryGatewayRequestError, string(request)))
			newSignal.SetError(maskAny(invalidFactoryGatewayRequestError))
			responder <- newSignal
			continue
		}
		if err != nil {
			s.Log.V(3).Errorf("%#v", maskAny(err))
			newSignal.SetError(maskAny(err))
			responder <- newSignal
			continue
		}
		newSignal.SetObject("response", response)

		responder <- newSignal
	}
}

func (s *server) NewCore() (spec.Core, error) {
	s.Log.V(11).Debugf("call Factory.NewCore")

	var err error

	newConfig := core.DefaultConfig()
	newConfig.FactoryClient = s.FactoryClient
	newConfig.Log = s.Log
	newConfig.State, err = s.NewState(common.ObjectType.Core)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.TextGateway = s.TextGateway
	newCore := core.NewCore(newConfig)

	return newCore, nil
}

func (s *server) NewImpulse() (spec.Impulse, error) {
	s.Log.V(11).Debugf("call Factory.NewImpulse")

	var err error

	newConfig := impulse.DefaultConfig()
	newConfig.FactoryClient = s.FactoryClient
	newConfig.Log = s.Log
	newConfig.State, err = s.NewState(common.ObjectType.Impulse)
	if err != nil {
		return nil, maskAny(err)
	}
	newImpulse := impulse.NewImpulse(newConfig)

	return newImpulse, nil
}

func (s *server) NewCharacterNeuron() (spec.Neuron, error) {
	s.Log.V(11).Debugf("call Factory.NewCharacterNeuron")

	var err error

	newConfig := neuron.DefaultCharacterNeuronConfig()
	newConfig.FactoryClient = s.FactoryClient
	newConfig.Log = s.Log
	newConfig.State, err = s.NewState(common.ObjectType.CharacterNeuron)
	if err != nil {
		return nil, maskAny(err)
	}
	newNeuron := neuron.NewCharacterNeuron(newConfig)

	return newNeuron, nil
}

func (s *server) NewFirstNeuron() (spec.Neuron, error) {
	s.Log.V(11).Debugf("call Factory.NewFirstNeuron")

	var err error

	newConfig := neuron.DefaultFirstNeuronConfig()
	newConfig.FactoryClient = s.FactoryClient
	newConfig.Log = s.Log
	newConfig.State, err = s.NewState(common.ObjectType.FirstNeuron)
	if err != nil {
		return nil, maskAny(err)
	}
	newNeuron := neuron.NewFirstNeuron(newConfig)

	return newNeuron, nil
}

func (s *server) NewJobNeuron() (spec.Neuron, error) {
	s.Log.V(11).Debugf("call Factory.NewJobNeuron")

	var err error

	newConfig := neuron.DefaultJobNeuronConfig()
	newConfig.FactoryClient = s.FactoryClient
	newConfig.Log = s.Log
	newConfig.State, err = s.NewState(common.ObjectType.JobNeuron)
	if err != nil {
		return nil, maskAny(err)
	}
	newNeuron := neuron.NewJobNeuron(newConfig)

	return newNeuron, nil
}

func (s *server) NewNetwork() (spec.Network, error) {
	s.Log.V(11).Debugf("call Factory.NewNetwork")

	var err error

	newConfig := network.DefaultConfig()
	newConfig.FactoryClient = s.FactoryClient
	newConfig.Log = s.Log
	newConfig.State, err = s.NewState(common.ObjectType.Network)
	if err != nil {
		return nil, maskAny(err)
	}
	newNetwork := network.NewNetwork(newConfig)

	return newNetwork, nil
}

func (s *server) NewState(objectType spec.ObjectType) (spec.State, error) {
	s.Log.V(11).Debugf("call Factory.NewState")

	newStateConfig := state.DefaultConfig()
	newStateConfig.FactoryClient = s.FactoryClient
	newStateConfig.Log = s.Log
	newStateConfig.ObjectType = objectType
	newStateConfig.StateReader = s.StateReader
	newStateConfig.StateWriter = s.StateWriter
	newState := state.NewState(newStateConfig)

	return newState, nil
}
