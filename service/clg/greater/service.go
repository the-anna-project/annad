// Package greater implements spec.CLG and provides a method to identify which
// of the given numbers is greater than the other.
package greater

import (
	"github.com/the-anna-project/spec/object"
)

// calculate returns the number that is greater than the other.
func (s *service) calculate(ctx spec.Context, a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}
