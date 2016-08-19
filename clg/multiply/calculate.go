// Package multiply implements spec.CLG and provides the mathematical operation
// of multiplication.
package multiply

import (
	"golang.org/x/net/context"
)

// calculate creates the product of the given float64s.
func (c *clg) calculate(ctx context.Context, a, b float64) float64 {
	return a * b
}
