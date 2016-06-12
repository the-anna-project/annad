package clg

// IsBetween checks whether a given number lies between two given numbers.
func (c Collection) IsBetween(n, min, max float64) bool {
	if n < min {
		return false
	}
	if n > max {
		return false
	}
	return true
}

// Difference creates the difference of the given float64s.
func (c Collection) Difference(a, b float64) float64 {
	return a - b
}

// Greater returns the number that is greater than the other.
func (c Collection) Greater(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}

// IsGreater checks whether teh first given number is greater than the other.
func (c Collection) IsGreater(a, b float64) bool {
	return a > b
}

// Product creates the product of the given float64s.
func (c Collection) Product(a, b float64) float64 {
	return a * b
}

// Sum creates the sum of the given float64s.
func (c Collection) Sum(a, b float64) float64 {
	return a + b
}

// Quotient creates the quotient of the given float64s.
func (c Collection) Quotient(a, b float64) float64 {
	return a / b
}
