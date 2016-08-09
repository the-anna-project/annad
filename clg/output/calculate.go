package output

import (
	"github.com/xh3b4sd/anna/spec"
)

// TODO
func (n *network) calculate() error {
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
