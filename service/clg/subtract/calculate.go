// Package subtract implements spec.CLG and provides the mathematical operation
// of subtraction.
package subtract

import (
	"github.com/xh3b4sd/anna/object/spec"
)

// calculate creates the difference of the given float64s.
func (s *service) calculate(ctx spec.Context, a, b float64) float64 {
	return a - b
}
