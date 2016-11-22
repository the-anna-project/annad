package round

import (
	"testing"

	"github.com/the-anna-project/annad/object/context"
)

func Test_CLG_Round(t *testing.T) {
	testCases := []struct {
		Float     float64
		Precision int
		Expected  float64
	}{
		{
			Float:     3.5,
			Precision: 0,
			Expected:  4,
		},
		{
			Float:     3.4,
			Precision: 0,
			Expected:  3,
		},
		{
			Float:     3.4,
			Precision: 1,
			Expected:  3.4,
		},
		{
			Float:     3.4,
			Precision: 2,
			Expected:  3.4,
		},
		{
			Float:     3.476,
			Precision: 2,
			Expected:  3.48,
		},
		{
			Float:     -3.476,
			Precision: 2,
			Expected:  -3.48,
		},
		{
			Float:     3,
			Precision: 0,
			Expected:  3,
		},
		{
			Float:     3,
			Precision: 2,
			Expected:  3,
		},
	}

	newCLG := MustNew()

	for i, testCase := range testCases {
		f, err := newCLG.(*clg).calculate(context.MustNew(), testCase.Float, testCase.Precision)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		if f != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", f)
		}
	}
}

func Test_CLG_Round_Error_NegativePrecision(t *testing.T) {
	newCLG := MustNew()
	_, err := newCLG.(*clg).calculate(context.MustNew(), 3.4465, -3)
	if !IsParseFloatSyntax(err) {
		t.Fatal("case", "expected", true, "got", false)
	}
}
