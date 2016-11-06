// Package divide implements spec.CLG and provides the mathematical operation
// of division.
package divide

import (
	"github.com/xh3b4sd/anna/object/spec"
)

// calculate creates the quotient of the given float64s.
func (c *clg) calculate(ctx spec.Context, a, b float64) float64 {
	return a / b
}
