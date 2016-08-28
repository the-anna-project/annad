// Package multiply implements spec.CLG and provides the mathematical operation
// of multiplication.
package multiply

import (
	"github.com/xh3b4sd/anna/spec"
)

// calculate creates the product of the given float64s.
func (c *clg) calculate(ctx spec.Context, a, b float64) float64 {
	return a * b
}
