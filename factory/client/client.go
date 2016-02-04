package factoryclient

import (
	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/gateway/spec"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

type Config struct {
	FactoryGateway gatewayspec.Gateway

	Log spec.Log
}

func DefaultConfig() Config {
	config := Config{
		FactoryGateway: gateway.NewGateway(),
		Log:            log.NewLog(log.DefaultConfig()),
	}

	return config
}

func NewFactory(config Config) spec.Factory {
	newFactory := &client{
		Config: config,
	}

	return newFactory
}

type objectValues map[spec.ObjectID]spec.Object

type client struct {
	Config
}

func (c *client) NewCore() (spec.Core, error) {
	c.Log.V(11).Debugf("call FactoryClient.NewCore")

	bytes := map[string][]byte{
		"request": []byte(common.ObjectType.Core),
	}
	response, err := common.ForwardSignal(c.FactoryGateway, bytes)
	if err != nil {
		return nil, maskAny(err)
	}

	return response.(spec.Core), nil
}

func (c *client) NewImpulse() (spec.Impulse, error) {
	c.Log.V(11).Debugf("call FactoryClient.NewImpulse")

	bytes := map[string][]byte{
		"request": []byte(common.ObjectType.Impulse),
	}
	response, err := common.ForwardSignal(c.FactoryGateway, bytes)
	if err != nil {
		return nil, maskAny(err)
	}

	return response.(spec.Impulse), nil
}

func (c *client) NewCharacterNeuron() (spec.Neuron, error) {
	c.Log.V(11).Debugf("call FactoryClient.NewCharacterNeuron")

	bytes := map[string][]byte{
		"request": []byte(common.ObjectType.CharacterNeuron),
	}
	response, err := common.ForwardSignal(c.FactoryGateway, bytes)
	if err != nil {
		return nil, maskAny(err)
	}

	return response.(spec.Neuron), nil
}

func (c *client) NewFirstNeuron() (spec.Neuron, error) {
	c.Log.V(11).Debugf("call FactoryClient.NewFirstNeuron")

	bytes := map[string][]byte{
		"request": []byte(common.ObjectType.FirstNeuron),
	}
	response, err := common.ForwardSignal(c.FactoryGateway, bytes)
	if err != nil {
		return nil, maskAny(err)
	}

	return response.(spec.Neuron), nil
}

func (c *client) NewJobNeuron() (spec.Neuron, error) {
	c.Log.V(11).Debugf("call FactoryClient.NewFirstNeuron")

	bytes := map[string][]byte{
		"request": []byte(common.ObjectType.JobNeuron),
	}
	response, err := common.ForwardSignal(c.FactoryGateway, bytes)
	if err != nil {
		return nil, maskAny(err)
	}

	return response.(spec.Neuron), nil
}

func (c *client) NewNetwork() (spec.Network, error) {
	c.Log.V(11).Debugf("call FactoryClient.NewNetwork")

	bytes := map[string][]byte{
		"request": []byte(common.ObjectType.Network),
	}
	response, err := common.ForwardSignal(c.FactoryGateway, bytes)
	if err != nil {
		return nil, maskAny(err)
	}

	return response.(spec.Network), nil
}

func (c *client) NewState(objectType spec.ObjectType) (spec.State, error) {
	c.Log.V(11).Debugf("call FactoryClient.NewState")

	bytes := map[string][]byte{
		"state-object-type": []byte(objectType),
		"request":           []byte(common.ObjectType.State),
	}
	response, err := common.ForwardSignal(c.FactoryGateway, bytes)
	if err != nil {
		return nil, maskAny(err)
	}

	return response.(spec.State), nil
}
