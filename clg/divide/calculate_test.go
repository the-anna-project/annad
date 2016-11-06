package divide

import (
	"testing"

	"github.com/xh3b4sd/anna/object/context"
)

func Test_CLG_Divide(t *testing.T) {
	testCases := []struct {
		A        float64
		B        float64
		Expected float64
	}{
		{
			A:        3.5,
			B:        12.5,
			Expected: 0.28,
		},
		{
			A:        35.5,
			B:        14.5,
			Expected: 2.4482758620689653,
		},
		{
			A:        -3.5,
			B:        7.5,
			Expected: -0.4666666666666667,
		},
		{
			A:        12.5,
			B:        4.5,
			Expected: 2.7777777777777777,
		},
		{
			A:        36.5,
			B:        6.5,
			Expected: 5.615384615384615,
		},
	}

	newCLG := MustNew()

	for i, testCase := range testCases {
		f := newCLG.(*clg).calculate(context.MustNew(), testCase.A, testCase.B)
		if f != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", f)
		}
	}
}
