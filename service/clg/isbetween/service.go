// Package isbetween implements spec.CLG and provides a method to identify if a
// given number is between a given range.
package isbetween

import (
	objectspec "github.com/the-anna-project/spec/object"
)

// calculate checks whether a given number lies between two given numbers.
func (s *service) calculate(ctx objectspec.Context, n, min, max float64) bool {
	if n < min {
		return false
	}
	if n > max {
		return false
	}
	return true
}
