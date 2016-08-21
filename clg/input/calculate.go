// Package input implements spec.CLG and provides the entry to the neural
// network.
package input

import (
	"golang.org/x/net/context"
)

// calculate currently only returns the given input. TODO change this
func (c *clg) calculate(ctx context.Context, input string) string {
	return input
}
