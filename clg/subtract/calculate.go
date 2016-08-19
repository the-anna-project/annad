// Package subtract implements spec.CLG and provides the mathematical operation
// of subtraction.
package subtract

import (
	"golang.org/x/net/context"
)

// calculate creates the difference of the given float64s.
func (c *clg) calculate(ctx context.Context, a, b float64) float64 {
	return a - b
}
