// Package greater implements spec.CLG and provides a method to identify which
// of the given numbers is greater than the other.
package greater

import (
	"github.com/xh3b4sd/anna/object/spec"
)

// calculate returns the number that is greater than the other.
func (s *service) calculate(ctx spec.Context, a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}
