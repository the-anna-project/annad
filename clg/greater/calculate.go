package greater

// calculate returns the number that is greater than the other.
func (c *clg) calculate(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}
