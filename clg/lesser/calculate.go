// Package lesser implements spec.CLG and provides a method to identify which of
// the given numbers is lesser than the other.
package lesser

import (
	"github.com/xh3b4sd/anna/object/spec"
)

// calculate returns the number that is lesser than the other.
func (c *clg) calculate(ctx spec.Context, a, b float64) float64 {
	if a < b {
		return a
	}

	return b
}
