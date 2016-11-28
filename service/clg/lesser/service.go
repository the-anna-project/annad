// Package lesser implements spec.CLG and provides a method to identify which of
// the given numbers is lesser than the other.
package lesser

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// calculate returns the number that is lesser than the other.
func (s *service) calculate(ctx objectspec.Context, a, b float64) float64 {
	if a < b {
		return a
	}

	return b
}
