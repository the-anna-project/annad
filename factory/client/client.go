// FactoryClient implements spec.Factory and provides a decentralized service for
// general object creation.
package factoryclient

import (
	"sync"

	"github.com/xh3b4sd/anna/factory/common"
	"github.com/xh3b4sd/anna/gateway"
	"github.com/xh3b4sd/anna/id"
	"github.com/xh3b4sd/anna/log"
	"github.com/xh3b4sd/anna/spec"
)

const (
	ObjectTypeFactoryClient spec.ObjectType = "factory-client"
)

type Config struct {
	FactoryGateway spec.Gateway

	Log spec.Log
}

func DefaultConfig() Config {
	config := Config{
		FactoryGateway: gateway.NewGateway(gateway.DefaultConfig()),
		Log:            log.NewLog(log.DefaultConfig()),
	}

	return config
}

func NewFactory(config Config) spec.Factory {
	newFactory := &factoryClient{
		Closed: false,
		Closer: make(chan struct{}, 1),
		Config: config,
		ID:     id.NewObjectID(id.Hex128),
		Mutex:  sync.Mutex{},
		Type:   ObjectTypeFactoryClient,
	}

	newFactory.Log.Register(newFactory.GetType())

	return newFactory
}

type factoryClient struct {
	Closed bool
	Closer chan struct{}

	Config

	ID spec.ObjectID `json:"id"`

	Mutex sync.Mutex

	Type spec.ObjectType `json:"type"`
}

func (fc *factoryClient) Boot() {
	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call Boot")
}

func (fc *factoryClient) NewCore() (spec.Core, error) {
	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call NewCore")

	output, err := forwardSignal(fc.FactoryGateway, common.ObjectTypeCore, fc.Closer)
	if err != nil {
		return nil, maskAny(err)
	}

	return output.(spec.Core), nil
}

func (fc *factoryClient) NewImpulse() (spec.Impulse, error) {
	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call NewImpulse")

	output, err := forwardSignal(fc.FactoryGateway, common.ObjectTypeImpulse, fc.Closer)
	if err != nil {
		return nil, maskAny(err)
	}

	return output.(spec.Impulse), nil
}

func (fc *factoryClient) NewRedisStorage() (spec.Storage, error) {
	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call NewRedisStorage")

	output, err := forwardSignal(fc.FactoryGateway, common.ObjectTypeRedisStorage, fc.Closer)
	if err != nil {
		return nil, maskAny(err)
	}

	return output.(spec.Storage), nil
}

func (fc *factoryClient) NewStrategyNetwork() (spec.Network, error) {
	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call NewNetwork")

	output, err := forwardSignal(fc.FactoryGateway, common.ObjectTypeStrategyNetwork, fc.Closer)
	if err != nil {
		return nil, maskAny(err)
	}

	return output.(spec.Network), nil
}

func (fc *factoryClient) Shutdown() {
	fc.Log.WithTags(spec.Tags{L: "D", O: fc, T: nil, V: 15}, "call Shutdown")

	fc.Mutex.Lock()
	defer fc.Mutex.Unlock()

	fc.FactoryGateway.Close()

	if !fc.Closed {
		fc.Closer <- struct{}{}
		fc.Closed = true
	}
}
