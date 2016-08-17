// Package sum implements spec.CLG and provides the mathematical operation of
// addition.
package sum

import (
	"golang.org/x/net/context"
)

// calculate creates the sum of the given float64s.
func (c *clg) calculate(ctx context.Context, a, b float64) (context.Context, float64) {
	return ctx, a + b
}
