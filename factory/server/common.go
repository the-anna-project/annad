package factoryserver

import (
	"github.com/xh3b4sd/anna/factory/common"
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
	case common.ObjectTypeImpulse:
		output, err = fs.NewImpulse()
	default:
		return nil, maskAny(invalidFactoryGatewayRequestError)
	}

	if err != nil {
		return nil, maskAny(err)
	}

	newSignal.SetOutput(output)

	return newSignal, nil
}
