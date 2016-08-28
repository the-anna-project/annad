// Package isbetween implements spec.CLG and provides a method to identify if a
// given number is between a given range.
package isbetween

import (
	"github.com/xh3b4sd/anna/spec"
)

// calculate checks whether a given number lies between two given numbers.
func (c *clg) calculate(ctx spec.Context, n, min, max float64) bool {
	if n < min {
		return false
	}
	if n > max {
		return false
	}
	return true
}
