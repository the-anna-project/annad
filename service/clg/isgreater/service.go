// Package isgreater implements spec.CLG and provides a method to identify if
// the first given number is greater than the later.
package isgreater

import (
	"github.com/xh3b4sd/anna/object/spec"
)

// calculate checks whether the first given number is greater than the other.
func (s *service) calculate(ctx spec.Context, a, b float64) bool {
	return a > b
}
