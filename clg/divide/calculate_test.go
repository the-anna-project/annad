package divide

import (
	"reflect"
	"testing"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
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
	ctx := context.Background()

	for i, testCase := range testCases {
		newNetworkPayloadConfig := api.DefaultNetworkPayloadConfig()
		newNetworkPayloadConfig.Args = []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(testCase.A), reflect.ValueOf(testCase.B)}
		newNetworkPayloadConfig.Destination = "destination"
		newNetworkPayloadConfig.Sources = []spec.ObjectID{"source"}
		newNetworkPayload, err := api.NewNetworkPayload(newNetworkPayloadConfig)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}

		calculatedNetworkPayload, err := newCLG.Calculate(newNetworkPayload)
		if err != nil {
			t.Fatal("case", i+1, "expected", nil, "got", err)
		}
		args := calculatedNetworkPayload.GetArgs()
		if len(args) != 2 {
			t.Fatal("case", i+1, "expected", 2, "got", len(args))
		}
		switch ctx := args[0].Interface().(type) {
		case context.Context:
			// all good
		default:
			t.Fatal("case", i+1, "expected", "context.Context", "got", ctx)
		}
		result := args[1].Float()

		if result != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", result)
		}
	}
}
