// Package divide implements spec.CLG and provides the mathematical operation
// of division.
package divide

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// calculate creates the quotient of the given float64s.
func (s *service) calculate(ctx objectspec.Context, a, b float64) float64 {
	return a / b
}
