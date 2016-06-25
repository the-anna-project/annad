package core

import (
	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/expectation"
	"github.com/xh3b4sd/anna/impulse"
	"github.com/xh3b4sd/anna/spec"
)

func (n *network) NewImpulse(coreRequest api.CoreRequest) (spec.Impulse, error) {
	n.Log.WithTags(spec.Tags{L: "D", O: n, T: nil, V: 15}, "call NewImpulse")

	// Create a new expectation based on the given request.
	newExpectationConfig := expectation.DefaultConfig()
	newExpectationConfig.ExpectationRequest = coreRequest.Expectation
	newExpectation, err := expectation.New(newExpectationConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	// Create a new impulse based on the given input and expectation.
	newConfig := impulse.DefaultConfig()
	newConfig.Expectation = newExpectation
	newConfig.Input = coreRequest.Input
	newConfig.Log = n.Log
	newImpulse, err := impulse.New(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	return newImpulse, nil
}
