// Package readseparator implements spec.CLG and provides functionality to
// read a separator stored in association to a specific behavior ID.
package readseparator

import (
	"encoding/json"

	"github.com/xh3b4sd/anna/index/clg/collection/feature-set"
	"github.com/xh3b4sd/anna/key"
	"github.com/xh3b4sd/anna/spec"
)

func (c *clg) calculate(ctx spec.Context) (separator, error) {

	// TODO check if there is already a configuration for the own behavior ID
	// TODO if there is a configuration, read the separator
	// TODO if there is no configuration, read some sequence and store it as configuration
	// TODO return separator

	randomKey, err := c.Storage().Feature().GetRandomKey()
	if err != nil {
		return "", maskAny(err)
	}

	var separator string

	return separator, nil
}
