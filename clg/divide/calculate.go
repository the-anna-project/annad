// Package divide implements spec.CLG and provides the mathematical operation
// of division.
package divide

// calculate creates the quotient of the given float64s.
func (c *clg) calculate(a, b float64) float64 {
	return a / b
}
