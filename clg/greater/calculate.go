// Package greater implements spec.CLG and provides a method to identify which
// of the given numbers is greater than the other.
package greater

import (
	"golang.org/x/net/context"
)

// calculate returns the number that is greater than the other.
func (c *clg) calculate(ctx context.Context, a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}
