package factoryserver

import (
	"github.com/xh3b4sd/anna/common"
	"github.com/xh3b4sd/anna/spec"
)

func (fs *factoryServer) gatewayListener(newSignal spec.Signal) (spec.Signal, error) {
	input := newSignal.GetInput()

	var output interface{}
	var err error

	if input == nil {
		return nil, maskAny(invalidFactoryGatewayRequestError)
	}

	switch input.(spec.ObjectType) {
	case common.ObjectType.Core:
		output, err = fs.NewCore()
	case common.ObjectType.Impulse:
		output, err = fs.NewImpulse()
	case common.ObjectType.RedisStorage:
		output, err = fs.NewRedisStorage()
	case common.ObjectType.StrategyNetwork:
		output, err = fs.NewStrategyNetwork()
	default:
		return nil, maskAny(invalidFactoryGatewayRequestError)
	}

	if err != nil {
		return nil, maskAny(err)
	}

	newSignal.SetOutput(output)

	return newSignal, nil
}
