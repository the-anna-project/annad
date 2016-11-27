// Package multiply implements spec.CLG and provides the mathematical operation
// of multiplication.
package multiply

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// calculate creates the product of the given float64s.
func (s *service) calculate(ctx objectspec.Context, a, b float64) float64 {
	return a * b
}
