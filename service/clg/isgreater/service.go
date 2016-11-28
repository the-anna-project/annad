// Package isgreater implements spec.CLG and provides a method to identify if
// the first given number is greater than the later.
package isgreater

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// calculate checks whether the first given number is greater than the other.
func (s *service) calculate(ctx objectspec.Context, a, b float64) bool {
	return a > b
}
