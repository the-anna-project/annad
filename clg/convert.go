package clg

import (
	"strconv"
)

func (i *clgIndex) IntStringConvert(args ...interface{}) ([]interface{}, error) {
	n, err := ArgToInt(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	newString := strconv.Itoa(n)

	return []interface{}{newString}, nil
}
