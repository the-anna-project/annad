// Package isbetween implements spec.CLG and provides a method to identify if a
// given number is between a given range.
package isbetween

import (
	"golang.org/x/net/context"
)

// calculate checks whether a given number lies between two given numbers.
func (c *clg) calculate(ctx context.Context, n, min, max float64) bool {
	if n < min {
		return false
	}
	if n > max {
		return false
	}
	return true
}
