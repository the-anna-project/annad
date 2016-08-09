package isbetween

// calculate checks whether a given number lies between two given numbers.
func (c *clg) calculate(n, min, max float64) bool {
	if n < min {
		return false
	}
	if n > max {
		return false
	}
	return true
}
