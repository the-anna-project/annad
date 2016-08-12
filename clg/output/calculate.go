// Package output implements spec.CLG and provides one of the two very special
// CLGs. That is the output CLG. Its purpose is to check if the calculated
// output matches the provided expectation, if any expectation given. The
// output CLG is handled in a special way because it determines the end of all
// requested calculations within the neural network. After the output CLG has
// been executed, the calculated output is returned back to the requesting
// client.
package output

import (
	"github.com/xh3b4sd/anna/spec"
)

// TODO
func (c *clg) calculate() error {
	var e spec.Expectation

	// Check the calculated output aganst the provided expectation, if any. In
	// case there is no expectation provided, we simply go with what we
	// calculated. This then means we are probably not in a training situation.
	if e.IsEmpty() {
		return nil
	}

	// There is an expectation provided. Thus we are going to check the
	// calculated output against it. In case the provided expectation did match
	// the calculated result, we simply return it and stop the iteration.
	match, err := e.Match()
	if err != nil {
		return maskAny(err)
	}
	if match {
		return nil
	}

	// TODO move reward/punish to output CLG?
	// TODO expectation met == reward?
	// TODO expectation NOT met == punishing?

	return nil
}
