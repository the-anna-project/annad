// Package sum implements spec.CLG and provides the mathematical operation of
// addition.
package sum

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// calculate creates the sum of the given float64s.
func (s *service) calculate(ctx objectspec.Context, a, b float64) float64 {
	return a + b
}
