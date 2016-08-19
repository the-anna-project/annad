// Package divide implements spec.CLG and provides the mathematical operation
// of division.
package divide

import (
	"golang.org/x/net/context"
)

// calculate creates the quotient of the given float64s.
func (c *clg) calculate(ctx context.Context, a, b float64) float64 {
	return a / b
}
