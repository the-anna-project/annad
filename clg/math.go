package clg

// Between checks whether a given number lies between two given numbers.
func (c Collection) Between(n, min, max float64) bool {
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
