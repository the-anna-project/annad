// Package round implements spec.CLG and provides a method to round the given
// number using the given precision.
package round

import (
	"fmt"
	"strconv"

	"github.com/xh3b4sd/anna/object/spec"
)

func (c *clg) calculate(ctx spec.Context, f float64, p int) (float64, error) {
	rounded, err := strconv.ParseFloat(fmt.Sprintf(fmt.Sprintf("%%.%df", p), f), 64)
	if err != nil {
		return 0, maskAny(err)
	}

	return rounded, nil
}
