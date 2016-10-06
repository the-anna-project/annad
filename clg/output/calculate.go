// Package output implements spec.CLG and provides one of the two very special
// CLGs. That is the output CLG. Its purpose is to check if the calculated
// output matches the provided expectation, if any expectation given. The
// output CLG is handled in a special way because it determines the end of all
// requested calculations within the neural network. After the output CLG has
// been executed, the calculated output is returned back to the requesting
// client.
package output

import (
	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

// TODO there is no CLG to read from the certenty pyramid

func (c *clg) calculate(ctx spec.Context, informationSequence string) error {
	// Check the calculated output against the provided expectation, if any. In
	// case there is no expectation provided, we simply go with what we
	// calculated. This then means we are probably not in a training situation.
	expectation, ok := ctx.GetExpectation()
	if !ok {
		err := c.sendTextResponse(informationSequence)
		if err != nil {
			return maskAny(err)
		}

		return nil
	}

	// There is an expectation provided. Thus we are going to check the calculated
	// output against it. In case the provided expectation does match the
	// calculated result, we simply return it.
	calculatedOutput := expectation.GetOutput()
	if informationSequence == calculatedOutput {
		err := c.sendTextResponse(informationSequence)
		if err != nil {
			return maskAny(err)
		}
	}

	return maskAnyf(expectationNotMetError, "'%s' != '%s'", informationSequence, calculatedOutput)
}

func (c *clg) sendTextResponse(informationSequence string) error {
	// Return the calculated output to the requesting client, if the
	// current CLG is the output CLG.
	newTextResponseConfig := api.DefaultTextResponseConfig()
	newTextResponseConfig.Output = informationSequence
	newTextResponse, err := api.NewTextResponse(newTextResponseConfig)
	if err != nil {
		return maskAny(err)
	}

	c.Gateway().TextOutput().GetChannel() <- newTextResponse

	return nil
}
