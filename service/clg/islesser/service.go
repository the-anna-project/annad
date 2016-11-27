// Package islesser implements spec.CLG and provides a method to identify if
// the first given number is lesser than the later.
package islesser

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// calculate checks whether the first given number is lesser than the other.
func (s *service) calculate(ctx objectspec.Context, a, b float64) bool {
	return a < b
}
