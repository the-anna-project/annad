package multiply

import (
	"reflect"
	"testing"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/context"
	"github.com/xh3b4sd/anna/spec"
)

func Test_CLG_Multiply(t *testing.T) {
	testCases := []struct {
		A        float64
		B        float64
		Expected float64
	}{
		{
			A:        3.5,
			B:        12.5,
			Expected: 43.75,
		},
		{
			A:        35.5,
			B:        14.5,
			Expected: 514.75,
		},
		{
			A:        -3.5,
			B:        7.5,
			Expected: -26.25,
		},
		{
			A:        12.5,
			B:        4.5,
			Expected: 56.25,
		},
		{
			A:        36.5,
			B:        6.5,
			Expected: 237.25,
		},
		{
			A:        17,
			B:        65,
			Expected: 1105,
		},
	}

	newCLG := MustNew()
	ctx := context.MustNew()

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
		result := args[1].Float()

		if result != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", result)
		}
	}
}
