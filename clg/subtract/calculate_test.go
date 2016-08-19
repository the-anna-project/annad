package subtract

import (
	"reflect"
	"testing"

	"golang.org/x/net/context"

	"github.com/xh3b4sd/anna/api"
	"github.com/xh3b4sd/anna/spec"
)

func Test_CLG_Subtract(t *testing.T) {
	testCases := []struct {
		A        float64
		B        float64
		Expected float64
	}{
		{
			A:        3.5,
			B:        12.5,
			Expected: -9,
		},
		{
			A:        35.5,
			B:        14.5,
			Expected: 21,
		},
		{
			A:        -3.5,
			B:        -7.5,
			Expected: 4,
		},
		{
			A:        12.5,
			B:        4.5,
			Expected: 8,
		},
		{
			A:        36.5,
			B:        6.5,
			Expected: 30,
		},
		{
			A:        11.11,
			B:        10.10,
			Expected: 1.0099999999999998,
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
		if len(args) != 1 {
			t.Fatal("case", i+1, "expected", 1, "got", len(args))
		}
		result := args[0].Float()

		if result != testCase.Expected {
			t.Fatal("case", i+1, "expected", testCase.Expected, "got", result)
		}
	}
}
